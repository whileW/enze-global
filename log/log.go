package log

import (
	"fmt"
	"go.uber.org/zap"
)

var log Log

func init()  {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("can't initialize zap logger: %v", err)
		return
	}
	log.log = logger.Sugar()
	defer log.log.Sync()
}

func InitLog() *Log {
	return &log
}