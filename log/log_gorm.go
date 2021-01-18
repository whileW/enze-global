package log

import (
	"context"
	"gorm.io/gorm/logger"
	"time"
)

// GormLogger struct
type GormLogger struct{
	LogLevel		logger.LogLevel
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		GetLoger().Infow(msg,data)
	}
}
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		GetLoger().Errorw(msg,data)
	}
}
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		GetLoger().Debugw(msg,data)
	}
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > logger.Silent {
		elapsed := time.Since(begin)
		sql, rows := fc()
		GetLoger().WithDuration(elapsed).Debugw("Gorm Trace",
			"sql",sql,
			"rows",rows)
	}
}

// Print - Log Formatter
func (*GormLogger) Print (v ...interface{}) {
	switch v[0] {
	case "sql":
		GetLoger().Debugw("sql",
			"module","gorm",
			"type","sql",
			"src",v[1],
			"duration",v[2],
			"sql",v[3],
			"values",v[4],
			"rows_returned",v[5])
	case "log":
		GetLoger().Debugw("gorm_log",
			"desc",v[2])
	}
}
