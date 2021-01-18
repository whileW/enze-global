package log

import (
	"go.uber.org/zap"
	"time"
)

type Loger struct {
	zap 		*zap.SugaredLogger

	duration	*time.Duration
}

func (l *Loger)WithDuration(duration time.Duration) *Loger {
	tlog := log
	tlog.duration = &duration
	return &tlog
}
func (l *Loger)handWith(args ...interface{}) []interface{} {
	if l.duration != nil {
		args = append(args, "duration",l.duration)
	}
	return args
}

func (l *Loger)Infow(msg string, keysAndValues ...interface{})  {
	keysAndValues = l.handWith(keysAndValues...)
	l.zap.Infow(msg,keysAndValues...)
}
func (l *Loger)Info(args ...interface{})  {
	args = l.handWith(args)
	l.zap.Info(args)
}
func (l *Loger)Errorw(msg string, keysAndValues ...interface{})  {
	keysAndValues = l.handWith(keysAndValues)
	l.zap.Errorw(msg,keysAndValues...)
}
func (l *Loger)Error(args ...interface{})  {
	args = l.handWith(args)
	l.zap.Error(args...)
}
func (l *Loger)Debugw(msg string, keysAndValues ...interface{})  {
	keysAndValues = l.handWith(keysAndValues)
	l.zap.Debugw(msg,keysAndValues...)
}
func (l *Loger)Debug(args ...interface{})  {
	args = l.handWith(args)
	l.zap.Debug(args...)
}