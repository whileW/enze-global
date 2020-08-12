package test

import (
	_"github.com/whileW/enze-utils/test/app"
	"fmt"
	"github.com/whileW/enze-utils"
	"testing"
)

func TestGetCurrentDirectory(t *testing.T)  {
	path := utils.GetCurrentDirectory()
	fmt.Println(path)
}