package test

import (
	"enze/global"
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	fmt.Println(global.GVA_CONFIG.SysSetting.Env)
}
