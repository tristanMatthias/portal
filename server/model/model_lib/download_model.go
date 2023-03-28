package model_lib

import (
	"fmt"
	"portal/server/lib"
)

func DownloadModel(hfId string, progress func(percentage float64)) {
	urls := getModelFiles(hfId)
	dest := lib.ModelPath(&hfId)

	progressChan, err := lib.DownloadFiles(dest, urls)
	if err != nil {
		fmt.Println(err)
		return
	}

	for p := range progressChan {
		progress(p)
	}
}

func getModelFiles(model string) []string {
	const baseURL = "https://huggingface.co"
	return []string{
		fmt.Sprintf("%s/%s/resolve/main/pytorch_model.bin", baseURL, model),
		fmt.Sprintf("%s/%s/resolve/main/config.json", baseURL, model),
		fmt.Sprintf("%s/%s/resolve/main/tokenizer_config.json", baseURL, model),
		fmt.Sprintf("%s/%s/resolve/main/vocab.json", baseURL, model),
		fmt.Sprintf("%s/%s/resolve/main/merges.txt", baseURL, model),
		fmt.Sprintf("%s/%s/resolve/main/vocab.txt", baseURL, model),
		fmt.Sprintf("%s/%s/resolve/main/special_tokens_map.json", baseURL, model),
		fmt.Sprintf("%s/%s/resolve/main/added_tokens.json", baseURL, model),
	}
}
