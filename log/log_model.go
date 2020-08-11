package log

import "go.uber.org/zap"

type Log struct {
	log 		*zap.SugaredLogger
}

func (l *Log)Infow(msg string, keysAndValues ...interface{})  {
	l.log.Infow(msg,keysAndValues...)
}
func (l *Log)Warnw(msg string, keysAndValues ...interface{})  {
	l.log.Warnw(msg,keysAndValues...)
}
func (l *Log)Debugw(msg string, keysAndValues ...interface{})  {
	l.log.Debugw(msg,keysAndValues...)
}