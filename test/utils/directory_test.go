package test

import (
	"github.com/whileW/enze-global/utils"
	_"github.com/whileW/enze-global/test/app"
	"fmt"
	"testing"
	"time"
)

func TestGetCurrentDirectory(t *testing.T)  {
	path := utils.GetCurrentDirectory()
	fmt.Println(path)
}

func TestCreateDir(t *testing.T)  {
	err := utils.CreateDir("upload/"+time.Now().Format("20060102"))
	fmt.Println(err)
}