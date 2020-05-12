package tools

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var (
	urlStr   = "https://www.baidu.com/img/bd_logo1.png"
	filename = fmt.Sprintf("/tmp/download_file_%d", time.Now().UnixNano())
)

func TestDownloadFile(t *testing.T) {
	err := DownloadFile(urlStr, 3*time.Second, filename)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		err := os.Remove(filename)
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestDownloadFile2(t *testing.T) {
	err := DownloadFile(urlStr, 1*time.Millisecond, filename)
	if err == nil {
		t.Fatalf("download %s expect timeout, actually not", urlStr)
	}
}

func TestDownloadFile3(t *testing.T) {
	err := DownloadFile("http://notexistdomain.com", 1*time.Second, filename)
	if err == nil {
		t.Fatalf("download %s expect error, actually not", urlStr)
	}
}

func TestDownloadFile4(t *testing.T) {
	err := DownloadFile(urlStr, 3*time.Second, "/not_exist_path/not_exist_file")
	if err == nil {
		t.Fatalf("download %s expect error, actually not", urlStr)
	}
}
