package initialize

import (
	"thingue-launcher/common/logger"
)

type ZapLogger struct{}

func (l *ZapLogger) Print(message string) {
	logger.Zap.Debug(message)
}

func (l *ZapLogger) Trace(message string) {
	logger.Zap.Debug(message)
}

func (l *ZapLogger) Debug(message string) {
	logger.Zap.Debug(message)
}

func (l *ZapLogger) Info(message string) {
	logger.Zap.Info(message)
}

func (l *ZapLogger) Warning(message string) {
	logger.Zap.Warn(message)
}

func (l *ZapLogger) Error(message string) {
	logger.Zap.Error(message)
}

func (l *ZapLogger) Fatal(message string) {
	logger.Zap.Fatal(message)
}
