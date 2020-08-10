package main

import (
	"github.com/whileW/enze-global"
	"fmt"
)

func main()  {
	fmt.Println(global.GVA_CONFIG.Setting)
	fmt.Println(global.GVA_CONFIG.Setting.GetChild("test").GetInt("test2"))
}