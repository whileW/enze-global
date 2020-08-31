package cache

import (
	"fmt"
	_"github.com/whileW/enze-global/test/app"
	"github.com/whileW/enze-global/utils/cache"
	"strconv"
	"testing"
)

func TestFifo(t *testing.T)  {
	f := cache.NewFifo()
	for i:=0 ;i<2000000 ;i++  {
		f.Push(strconv.Itoa(i),i)
	}
	fmt.Println(f)
}