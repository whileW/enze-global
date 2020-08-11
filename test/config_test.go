package test

import (
	_ "github.com/whileW/enze-global/test/app"
	"github.com/whileW/enze-global"
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	fmt.Println(global.GVA_CONFIG.SysSetting.Env)
}
