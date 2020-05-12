package tools

import (
	"io"
	"net/http"
	"os"
	"time"
)

func DownloadFile(urlStr string, timeout time.Duration, filename string) error {
	client := &http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
