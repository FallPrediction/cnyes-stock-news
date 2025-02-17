package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func getLoggerWriter() zapcore.WriteSyncer {
	os.Mkdir("./logs", os.ModePerm)
	file, _ := os.Create("./logs/cnyes-stock-news.log")
	return zapcore.AddSync(file)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func init() {
	core := zapcore.NewCore(getEncoder(), getLoggerWriter(), zap.ErrorLevel)
	Logger = zap.New(core).Sugar()
}
