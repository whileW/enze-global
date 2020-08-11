package log

import "go.uber.org/zap"

type Log struct {
	log 		*zap.SugaredLogger
}

func (l *Log)Infow(msg string, keysAndValues ...interface{})  {
	l.log.Infow(msg,keysAndValues...)
}
func (l *Log)Info(args ...interface{})  {
	l.log.Info(args...)
}
func (l *Log)Errorw(msg string, keysAndValues ...interface{})  {
	l.log.Errorw(msg,keysAndValues...)
}
func (l *Log)Error(args ...interface{})  {
	l.log.Error(args...)
}
func (l *Log)Debugw(msg string, keysAndValues ...interface{})  {
	l.log.Debugw(msg,keysAndValues...)
}
func (l *Log)Debug(args ...interface{})  {
	l.log.Debug(args...)
}