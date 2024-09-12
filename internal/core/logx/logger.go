package logx

import (
	"gin-quickly-template/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjackV2 "gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

func NameSpace(name string) *zap.SugaredLogger { return zap.S().Named(name) }

func getLogWriter() zapcore.WriteSyncer {
	if config.GetConfig().Log.LogPath == "" {
		log.Fatalln("LogPath isn't set")
	}
	lj := &lumberjackV2.Logger{
		Filename:   config.GetConfig().Log.LogPath,
		MaxSize:    5,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
	return zapcore.AddSync(lj)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func Init(level zapcore.LevelEnabler) {
	writeSyncer := getLogWriter()
	writeSyncer = zapcore.NewMultiWriteSyncer(writeSyncer, os.Stdout)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, level)
	options := []zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	}
	logger := zap.New(core, options...)

	// TODO: Add CLS hook
	//if CLSConfig := conf.Get().Log.CLS; CLSConfig.Endpoint != "" {
	//
	//}

	zap.ReplaceGlobals(logger)
}
