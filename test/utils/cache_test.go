package test

import (
	_ "github.com/whileW/enze-global/test/app"
	"github.com/whileW/enze-global/utils/cache"
	"strconv"
	"testing"
)

func TestFIFO(t *testing.T)  {
	f := cache.NewFIFO()
	for i:=0 ;i<2000000 ;i++  {
		f.Push(strconv.Itoa(i),i)
	}
}

