package test

import (
	//_ "github.com/whileW/enze-global/test/app"
	"github.com/whileW/enze-global"
	"fmt"
	"github.com/whileW/enze-global/config"
	"github.com/whileW/enze-global/initialize"
	"os"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
	os.Chdir("../runtest")
	config.InitConfg()
	initialize.Db()
	fmt.Println(global.GVA_CONFIG.SysSetting.Env)
	time.Sleep(500*time.Second)
}
