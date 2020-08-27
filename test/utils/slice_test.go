package test

import (
	"github.com/whileW/enze-global/utils"
	"testing"
)

func TestSliceStringContains(t *testing.T)  {
	a := []string{"1","2","3"}
	b := "2"
	t.Log(utils.SliceStringContains(a,b))
}