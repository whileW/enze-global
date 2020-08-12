package log

import (
	"github.com/whileW/enze-global/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var log Log

func init()  {
	filepath := getFilePath()
	logLevel := zap.NewAtomicLevel()

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  filepath,
		MaxSize:   1024, //MB
		LocalTime: true,
		Compress:  true,
	})

	allCore := []zapcore.Core{}
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	allCore = append(allCore, zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		//consoleEncoder,
		w,
		logLevel,
	))

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleDebugging := zapcore.Lock(os.Stdout)
	allCore = append(allCore, zapcore.NewCore(consoleEncoder, consoleDebugging, logLevel))

	core := zapcore.NewTee(allCore...)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	log.log = logger.Sugar()
	defer log.log.Sync()
}

func getFilePath() string {
	path := utils.GetCurrentDirectory()
	err := utils.CreateDir(path+ "/Log")
	if err == nil {
		path += "/Log"
	}
	logfile := path + "/" + utils.GetAppname() + ".log"
	return logfile
}

func InitLog() *Log {
	return &log
}