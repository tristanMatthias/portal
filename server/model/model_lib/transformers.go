package model_lib

import (
	"fmt"
)

const baseURL = "https://huggingface.co"

func ModelURL(model string) string {
	return fmt.Sprintf("%s/%s/resolve/main/pytorch_model.bin", baseURL, model)
}

func ConfigURL(model string) string {
	return fmt.Sprintf("%s/%s/resolve/main/config.json", baseURL, model)
}

func TokenizerURLs(model string) map[string]string {
	return map[string]string{
		"tokenizer_config.json": fmt.Sprintf("%s/%s/resolve/main/tokenizer_config.json", baseURL, model),
		"vocab.json":            fmt.Sprintf("%s/%s/resolve/main/vocab.json", baseURL, model),
		"merges.txt":            fmt.Sprintf("%s/%s/resolve/main/merges.txt", baseURL, model),
		"vocab.txt":             fmt.Sprintf("%s/%s/resolve/main/vocab.txt", baseURL, model),
		"special_tokens_map.json": fmt.Sprintf("%s/%s/resolve/main/special_tokens_map.json", baseURL, model),
		"added_tokens.json":     fmt.Sprintf("%s/%s/resolve/main/added_tokens.json", baseURL, model),
	}
}
