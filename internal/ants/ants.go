package ants

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
)

var W *ants.Pool

func init() {
	var err error
	W, err = ants.NewPool(10, ants.WithNonblocking(true))
	if err != nil {
		panic(fmt.Sprintln("New ants pool failed,err:", err))
	}
}
