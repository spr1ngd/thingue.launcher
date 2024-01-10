package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/provider"
)

var Zap *zap.SugaredLogger

func InitZapLogger() {
	level, err := zapcore.ParseLevel(provider.AppConfig.SystemSettings.LogLevel)
	if err != nil {
		fmt.Println("日志级别设置失败", err)
		level = zapcore.InfoLevel
	}
	fileCore := zapcore.NewCore(getFileEncoder(), getLogFileWriter(), level)
	consoleCore := zapcore.NewCore(getConsoleEncoder(), getLogConsoleWriter(), level)
	tee := zapcore.NewTee(fileCore, consoleCore)
	logger := zap.New(tee, zap.AddCaller())
	Zap = logger.Sugar()
}

func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getFileEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
	//return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	//return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogFileWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   constants.SAVE_DIR + "app.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getLogConsoleWriter() zapcore.WriteSyncer {
	return zapcore.Lock(os.Stdout)
}
