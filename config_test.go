package una

import (
	"os"
	"syscall"
	"testing"
	"time"
)

type AppConfig struct {
	Project string `toml:"project"`
	Num     uint   `toml:"num"`
}

func TestSetupConfig(t *testing.T) {
	tomlFilename := "./testdata/config_test.toml"
	var appConf AppConfig

	defer func() {
		if err := recover(); err != nil {
			t.Fatal(err)
		}
	}()

	SetupConfig(tomlFilename, &appConf)

	if appConf.Project != "ProjectName" {
		t.Fatalf("parse config project error, expect: %s, actual: %s", "ProjectName", appConf.Project)
	}
	if appConf.Num != 3 {
		t.Fatalf("parse config num error, expect: %d, actual: %d", 3, appConf.Num)
	}
}

func TestReloadConfigHandler(t *testing.T) {
	tomlFilename := "/tmp/config_test.toml"
	var appConf AppConfig

	f, err := os.OpenFile(tomlFilename, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	err = f.Truncate(0)
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString(`project="ProjectName"
num=3`)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := recover(); err != nil {
			t.Fatal(err)
		}
	}()

	SetupConfig(tomlFilename, &appConf)
	if appConf.Num != 3 {
		t.Fatalf("parse config num error, expect: %d, actual: %d", 3, appConf.Num)
	}

	ReloadConfigHandler()
	err = f.Truncate(0)
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString(`project="ProjectName"
num=4`)
	if err != nil {
		t.Fatal(err)
	}
	err = syscall.Kill(os.Getpid(), syscall.SIGHUP)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(5 * time.Millisecond)
	if appConf.Num != 4 {
		t.Fatalf("reload config num error, expect: %d, actual: %d", 4, appConf.Num)
	}
}
