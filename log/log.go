package log

import (
	"github.com/whileW/enze-global/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var log Loger

func GetLoger() *Loger {
	return &log
}

func init()  {
	filepath := getFilePath()
	logLevel := zap.NewAtomicLevel()

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  filepath,
		MaxSize:   10, //MB
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
	log.zap = logger.Sugar()
	defer log.zap.Sync()
}

func getFilePath() string {
	path := utils.GetCurrentDirectory() + "/Log"
	err := utils.CreateDir(path)
	if err != nil {
		panic("创建日志文件夹失败："+err.Error())
	}
	logfile := path + "/" + utils.GetAppname() + ".log"
	return logfile
}

func InitLog() *Loger {
	return &log
}