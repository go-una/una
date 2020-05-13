package una

import (
	"fmt"
	"os"

	"github.com/go-una/una/tools"
)

var (
	ProjectName = ""
	ModuleName  = ""
	RootPath    = ""
	ConfigPath  = ""
	LogsPath    = ""
	AssetsPath  = ""
	RunPath     = ""
	TmpPath     = ""
)

func Setup(options ProjectOptions) {
	ProjectName = options.Name
	ModuleName = options.Module
	RootPath = options.RootPath

	isRootPathExist, isRootPathDir := tools.FileExists(RootPath)
	if !isRootPathExist {
		panic(fmt.Sprintf("RootPath (%s) is not exist", RootPath))
	}
	if !isRootPathDir {
		panic(fmt.Sprintf("RootPath (%s) is not a directory", RootPath))
	}

	// env
	if options.Env == "" {
		options.Env = "prod"
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

	if options.LoggerOptions != nil {
		setupLogger(options.LoggerOptions)
	}
	if options.AccessLoggerOptions != nil {
		setupAccessLogger(options.AccessLoggerOptions)
	}

	configFilename := options.ConfigFilename
	if configFilename == "" {
		configFilename = fmt.Sprintf("%s/app.%s.toml", ConfigPath, options.Env)
	}
	SetupConfig(configFilename, options.Config)
}
