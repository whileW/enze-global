package log

import (
	"testing"
	"time"
)

func TestInitLog(t *testing.T) {
	GetLoger().Info("测试日志")
}

func TestLog_WithDuration(t *testing.T) {
	GetLoger().WithDuration(500*time.Second).Infow("测试日志","test","test")
	GetLoger().WithDuration(500*time.Second).Info("测试日志 %s %s")
}