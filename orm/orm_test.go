package orm

import (
	"fmt"
	"github.com/whileW/enze-global/config"
	"testing"
)

func TestInitOrm(t *testing.T) {
	config := config.InitConfg()
	orms := InitOrm(config)
	fmt.Println(orms.dbs)
}