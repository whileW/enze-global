package config

import (
	"fmt"
	"testing"
)

func TestInitConfg(t *testing.T) {
	conf := InitConfg()
	fmt.Println(conf)
}