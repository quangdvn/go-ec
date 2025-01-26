package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 1.
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name: %s, age: %d", "quangdvn", 25)

	// // Logger
	// logger := zap.NewExample()
	// logger.Info("Hello name: quangdvn", zap.String("name", "quangdvn"), zap.Int("age", 25))

	// 2.
	// logger := zap.NewExample()
	// logger.Info("Hello World")

	// // Dev
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello Dev")

	// // Prod
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello Prod")

	// 3.
	encoder := getLogEncoder()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))

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

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./logs/test_log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stdout)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
