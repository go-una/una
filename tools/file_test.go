package tools

import "testing"

func TestFileExists(t *testing.T) {
	filename := "./file.go"
	isExist, isDir := FileExists(filename)
	if !isExist {
		t.Fatalf("filename (%s) expect exist, actually not", filename)
	}
	if isDir {
		t.Fatalf("filename (%s) expect a file, actually is a directory", filename)
	}
}

func TestFileExists2(t *testing.T) {
	filepath := "../tools"
	isExist, isDir := FileExists(filepath)
	if !isExist {
		t.Fatalf("filepath (%s) expect exist, actually not", filepath)
	}
	if !isDir {
		t.Fatalf("filepath (%s) expect a directory, actually not", filepath)
	}
}

func TestFileExists3(t *testing.T) {
	filename := "./not_exist_file"
	isExist, isDir := FileExists(filename)
	if isExist {
		t.Fatalf("filename (%s) expect not exist, actually exist", filename)
	}
	if isDir {
		t.Fatalf("filename (%s) expect not a directory", filename)
	}
}
