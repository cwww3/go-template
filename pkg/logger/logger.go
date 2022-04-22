package logger

import (
	"github.com/cwww3/go-template/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	logger *zap.SugaredLogger
)

func GetLogger() *zap.SugaredLogger {
	return logger
}

func Init() {
	wss := make([]zapcore.WriteSyncer, 0, 2)
	zapConfig := config.GetZapConfig()
	if zapConfig.FileStdout {
		w := &lumberjack.Logger{
			Filename:   zapConfig.Path,
			MaxSize:    zapConfig.MaxSize,
			MaxAge:     zapConfig.MaxAge,
			MaxBackups: zapConfig.MaxBackups,
			LocalTime:  zapConfig.LocalTime,
			Compress:   zapConfig.Compress,
		}
		wss = append(wss, zapcore.AddSync(w))
	}
	if zapConfig.ConsoleStdout {
		wss = append(wss, zapcore.AddSync(os.Stdout))
	}

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(wss...), zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller()).Sugar()
}
