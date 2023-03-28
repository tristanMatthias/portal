package lib

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/cavaliergopher/grab/v3"
)

func DownloadFiles(dir string, urls []string) (chan float64, error) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	toDownload := []string{}
	totalSize := int64(0)

	for _, url := range urls {
		res, err := http.Head(url)
		if err != nil {
			fmt.Println(err)
		} else {
			// Add content-length to total size
			totalSize += int64(res.ContentLength)
			toDownload = append(toDownload, url)
		}
	}

	// Progress map for each file
	progress := make(map[string]int64)

	client := grab.NewClient()
	reqs := make([]*grab.Request, 0, len(toDownload))

	for _, url := range toDownload {
		// Set the filename to the last part of the URL
		dest := filepath.Join(dir, url[strings.LastIndex(url, "/")+1:])
		req, _ := grab.NewRequest(dest, url)
		reqs = append(reqs, req)
	}

	// Start downloads
	fmt.Printf("Downloading %d files...\n", len(reqs))
	wg := sync.WaitGroup{}

	progressChan := make(chan float64)

	for _, req := range reqs {
		fmt.Printf("Downloading %v...\n", req.URL())
		wg.Add(1)

		// Run each download in a goroutine and report progress every 200ms
		go func(req *grab.Request) {
			defer wg.Done()

			url := req.URL().String()

			resp := client.Do(req)
			t := time.NewTicker(200 * time.Millisecond)
			defer t.Stop()

			// An inlined function to update progress
			updateProgress := func() {
				// Calculate overall progress
				totalBytes := int64(0)
				for _, p := range progress {
					totalBytes += p
				}
				progressChan <- float64(totalBytes) / float64(totalSize)
			}

			for {
				select {
				case <-t.C:
					progress[url] = resp.BytesComplete()
					updateProgress()
				case <-resp.Done:
					progress[url] = resp.Size()
					updateProgress()
					return
				}
			}
		}(req)
	}

	go func() {
		wg.Wait()
		progressChan <- 1
		close(progressChan)
	}()

	return progressChan, nil
}
