package test

import (
	"github.com/whileW/enze-global"
	_ "github.com/whileW/enze-global/test/app"
	"testing"
)

func TestLog(t *testing.T)  {
	global.GVA_LOG.Infow("测试",
		"test","1",
		"test","2",
		"test","3")
}