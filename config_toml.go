package una

import (
	"io/ioutil"

	toml "github.com/pelletier/go-toml"
)

func Load(filename string) (*toml.Tree, error) {
	return toml.LoadFile(filename)
}

func Parse(filename string, v interface{}) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return toml.Unmarshal(content, v)
}
