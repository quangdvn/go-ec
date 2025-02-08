package loggers

import (
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/quangdvn/go-ec/pkg/settings"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	*zap.Logger
}

func NewLogger(config settings.LoggerSetting) *ZapLogger {
	// DEBUG >> INFO >> WARN >> ERROR >> FATAL >> PANIC
	basicLogLevel := config.LogLevel
	var level zapcore.Level
	switch basicLogLevel {
	case "DEBUG":
		level = zapcore.DebugLevel
	case "INFO":
		level = zapcore.InfoLevel
	case "WARN":
		level = zapcore.WarnLevel
	case "ERROR":
		level = zapcore.ErrorLevel
	case "FATAL":
		level = zapcore.FatalLevel
	case "PANIC":
		level = zapcore.PanicLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getLogEncoder()
	hook := lumberjack.Logger{
		Filename:   config.FileLogName,
		MaxSize:    config.MaxSize, // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,   // days
		Compress:   config.Compress, // disabled by default
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)

	return &ZapLogger{
		zap.New(
			core,
			zap.AddCaller(),
			zap.AddStacktrace(zap.ErrorLevel),
		),
	}

}

// Format log
func getLogEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   //1737901993.0755782 >> 2025-01-26T23:33:13.074+0900
	encoderConfig.TimeKey = "time"                          // ts >> Time
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // INFO >> INFO
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder // caller

	return zapcore.NewJSONEncoder(encoderConfig)
}
