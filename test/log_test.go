package test

import (
	_ "github.com/whileW/enze-global/test/app"
	"testing"
	"github.com/whileW/enze-global"
)

func TestLog(t *testing.T)  {
	global.GVA_LOG.Infow("测试",
		"test","1",
		"test","2",
		"test","3")
}