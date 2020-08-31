package test

import (
	"github.com/whileW/enze-global/utils"
	_"github.com/whileW/enze-global/test/app"
	"fmt"
	"io/ioutil"
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

func TestGetDirChile(t *testing.T)  {
	f,_ := ioutil.ReadDir("upload")
	for _,t := range f {
		fmt.Println(t.Name())
	}
}