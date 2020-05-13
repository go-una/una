package una

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	zapLogger    *zap.Logger
	accessLogger *zap.Logger
	sugarLogger  *zap.SugaredLogger
)

func setupLogger(options *LoggerOptions) {
	if zapLogger == nil {
		logFilenameWithoutExt := ProjectName
		if ModuleName != "" {
			logFilenameWithoutExt += "-" + ModuleName
		}
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s.log", LogsPath, logFilenameWithoutExt),
			MaxSize:    options.MaxSize, // MB
			MaxBackups: options.MaxBackups,
		})
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			w,
			zap.InfoLevel,
		)
		zapLogger = zap.New(core)
	}
}

func setupAccessLogger(options *LoggerOptions) {
	if accessLogger == nil {
		logFilenameWithoutExt := "access"
		if ModuleName != "" {
			logFilenameWithoutExt += "-" + ModuleName
		}
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s.log", LogsPath, logFilenameWithoutExt),
			MaxSize:    options.MaxSize, // MB
			MaxBackups: options.MaxBackups,
		})

		encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			MessageKey: "msg",
		})
		core := zapcore.NewCore(
			encoder,
			w,
			zap.InfoLevel,
		)
		accessLogger = zap.New(core)
	}
}

func Logger() *zap.Logger {
	return zapLogger
}

func AccessLogger() *zap.Logger {
	return accessLogger
}

func SugarLogger() *zap.SugaredLogger {
	if sugarLogger == nil {
		sugarLogger = zapLogger.Sugar()
	}
	return sugarLogger
}
