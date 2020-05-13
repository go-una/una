package una

import "testing"

var (
	tomlFilename     = "./testdata/config_test.toml"
	notExistFilename = "not_exist.toml"
)

func TestLoad(t *testing.T) {
	conf, err := Load(tomlFilename)
	if err != nil {
		t.Fatal(err)
	}
	if conf.Get("project").(string) != "ProjectName" {
		t.Fatalf("get project error (expected: %v, actual: %v)\n",
			"ProjectName",
			conf.Get("project").(string),
		)
	}
}

func TestLoad2(t *testing.T) {
	_, err := Load(notExistFilename)
	if err == nil {
		t.Fatal("load not exist filename must be error")
	}
}

type DbConfig struct {
	Host     string `toml:"host"`
	Port     uint   `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type TestConfig struct {
	Project string   `toml:"project"`
	Num     uint     `toml:"num"`
	Db      DbConfig `toml:"db"`
}

func TestParse(t *testing.T) {
	var conf TestConfig
	err := Parse(tomlFilename, &conf)
	if err != nil {
		t.Fatal(err)
	}
	if conf.Project != "ProjectName" {
		t.Fatalf("get project error (expected: %v, actual: %v)\n",
			"ProjectName",
			conf.Project,
		)
	}
	if conf.Num != 3 {
		t.Fatalf("get num error (expected: %v, actual: %v)\n",
			3,
			conf.Num,
		)
	}
	if conf.Db.Host != "127.0.0.1" {
		t.Fatalf("get db.host error (expected: %v, actual: %v)\n",
			"127.0.0.1",
			conf.Db.Host,
		)
	}
}

func TestParse2(t *testing.T) {
	var conf TestConfig
	err := Parse(notExistFilename, &conf)
	if err == nil {
		t.Fatal("parse not exist filename must be error")
	}
}
