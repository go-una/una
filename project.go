package una

import (
	"fmt"
	"os"

	"github.com/go-una/una/tools"
)

type ProjectOptions struct {
	Name     string
	RootPath string
}

var (
	ProjectName = ""
	RootPath    = ""
	ConfigPath  = ""
	LogsPath    = ""
	AssetsPath  = ""
	RunPath     = ""
	TmpPath     = ""
)

func Setup(options ProjectOptions) {
	ProjectName = options.Name
	RootPath = options.RootPath

	isRootPathExist, isRootPathDir := tools.FileExists(RootPath)
	if !isRootPathExist {
		panic(fmt.Sprintf("RootPath (%s) is not exist", RootPath))
	}
	if !isRootPathDir {
		panic(fmt.Sprintf("RootPath (%s) is not a directory", RootPath))
	}

	ConfigPath = RootPath + "/config"
	LogsPath = RootPath + "/logs"
	AssetsPath = RootPath + "/assets"
	RunPath = RootPath + "/run"
	TmpPath = RootPath + "/tmp"

	for _, filepath := range []string{ConfigPath, LogsPath, AssetsPath, RunPath, TmpPath} {
		isExist, isDir := tools.FileExists(filepath)
		if isExist && !isDir {
			panic(fmt.Sprintf("filepath (%s) is not a directory", filepath))
		} else if !isExist {
			err := os.MkdirAll(filepath, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}
}
