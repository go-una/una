package una

import (
	"os"
	"testing"

	"github.com/go-una/una/tools"
)

func TestSetup(t *testing.T) {
	rootPath := "/tmp/project_una"

	defer func() {
		err := recover()
		if err != nil {
			t.Fatal(err)
		}
	}()

	err := os.MkdirAll(rootPath, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(rootPath)

	type AppConfig struct {
		ProjectName string `toml:"project"`
		Num         int    `toml:"num"`
	}
	var appConf AppConfig
	Setup(ProjectOptions{
		Name:           "una",
		Module:         "api",
		RootPath:       rootPath,
		ConfigFilename: "./testdata/config_test.toml",
		Config:         &appConf,
	})

	isExist, isDir := tools.FileExists(ConfigPath)
	if !isExist || !isDir {
		t.Fatalf("ConfigPath (%s) is not a directory", ConfigPath)
	}
	isExist, isDir = tools.FileExists(RunPath)
	if !isExist || !isDir {
		t.Fatalf("RunPath (%s) is not a directory", RunPath)
	}
	isExist, isDir = tools.FileExists(TmpPath)
	if !isExist || !isDir {
		t.Fatalf("TmpPath (%s) is not a directory", TmpPath)
	}

	if ProjectName != "una" {
		t.Fatalf("ProjectName error, expect: una, actual: %s", ProjectName)
	}
	if ModuleName != "api" {
		t.Fatalf("ModuleName error, expect: api, actual: %s", ModuleName)
	}

	if appConf.ProjectName != "ProjectName" {
		t.Fatalf("parse app config error, project expect: %s, actual: %s", "ProjectName", appConf.ProjectName)
	}
	if appConf.Num != 3 {
		t.Fatalf("parse app config error, num expect: %d, actual: %d", 3, appConf.Num)
	}
}

func TestSetup2(t *testing.T) {
	rootPath := "/tmp/project_not_exist"

	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("rootPath (%s) expect not exist", rootPath)
		}
	}()

	Setup(ProjectOptions{
		Name:     "una",
		RootPath: rootPath,
	})
}

func TestSetup3(t *testing.T) {
	rootPath := "/tmp/project_no_permission"

	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("rootPath (%s) expect no permission", rootPath)
		}
	}()

	err := os.MkdirAll(rootPath, 0444)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(rootPath)
	Setup(ProjectOptions{
		Name:     "una",
		RootPath: rootPath,
	})
}

func TestSetup4(t *testing.T) {
	rootPath := "/tmp/project_file"

	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("rootPath (%s) expect direcotry", rootPath)
		}
	}()

	f, err := os.OpenFile(rootPath, os.O_CREATE, 0755)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(rootPath)
	defer f.Close()

	Setup(ProjectOptions{
		Name:     "una",
		RootPath: rootPath,
	})
}

func TestSetup5(t *testing.T) {
	rootPath := "/tmp/project_una"
	configPath := rootPath + "/config"

	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("configPath (%s) expect direcotry", configPath)
		}
	}()

	err := os.MkdirAll(rootPath, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(rootPath)

	f, err := os.OpenFile(configPath, os.O_CREATE, 0755)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(configPath)
	defer f.Close()

	Setup(ProjectOptions{
		Name:     "una",
		RootPath: rootPath,
	})
}
