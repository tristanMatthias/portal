package model

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type Download struct {
	URL      string
	Progress int64
	Total    int64
	Done     bool
}

var Downloads = make(map[string]*Download)
var DownloadsLock sync.Mutex

func DownloadModel(model string) {

	DownloadsLock.Lock()
	if _, ok := Downloads[model]; ok {
		DownloadsLock.Unlock()
		return
	}
	download := &Download{URL: model}
	Downloads[model] = download
	DownloadsLock.Unlock()

	modelPath := filepath.Join(os.Getenv("HOME"), ".portal", "models", model)
	err := os.MkdirAll(modelPath, 0755)
	if err != nil {
		fmt.Println("Mkdir error:", err)
		return
	}

	// Download model file
	modelFile := filepath.Join(modelPath, "pytorch_model.bin")
	downloadFile(modelFile, ModelURL(model), download)

	// Download config file
	configFile := filepath.Join(modelPath, "config.json")
	downloadFile(configFile, ConfigURL(model), download)

	// Download tokenizer files
	tokenizerFiles := TokenizerURLs(model)
	for filename, url := range tokenizerFiles {
		tokenizerFile := filepath.Join(modelPath, filename)
		downloadFile(tokenizerFile, url, download)
	}

	download.Done = true
}

func downloadFile(filepath, url string, download *Download) {
	out, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Create error:", err)
		return
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("File %s not found remotely\n", filepath)
		e := os.Remove(filepath)
		if e != nil {
			fmt.Println("Remove error:", e)
		}
		return
	}

	download.Total += resp.ContentLength
	_, err = io.Copy(out, &DownloadReader{download, resp.Body})
	if err != nil {
		fmt.Println("Copy error:", err)
		return
	}
}

type DownloadReader struct {
	download *Download
	reader   io.Reader
}

func (dr *DownloadReader) Read(p []byte) (int, error) {
	n, err := dr.reader.Read(p)
	dr.download.Progress += int64(n)
	return n, err
}
