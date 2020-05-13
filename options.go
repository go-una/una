package una

type ProjectOptions struct {
	Name           string
	Module         string
	RootPath       string
	Env            string      // 环境标识，dev: 开发环境；test：测试环境；prod：生产环境；
	ConfigFilename string      // 配置文件，若未指定，则默认为 {ConfigPath}/app.{Env}.toml
	Config         interface{} // 接收配置数据的参数

	LoggerOptions       *LoggerOptions
	AccessLoggerOptions *LoggerOptions
}

type LoggerOptions struct {
	MaxSize    int
	MaxBackups int
}
