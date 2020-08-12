package test

import (
	"github.com/whileW/enze-global/utils"
	_"github.com/whileW/enze-global/test/app"
	"fmt"
	"testing"
)

func TestGetCurrentDirectory(t *testing.T)  {
	path := utils.GetCurrentDirectory()
	fmt.Println(path)
}