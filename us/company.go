package us

import (
	"fmt"
	"math/rand"

	"../common"
)

var companySuffixes = []string{"Inc", "and Sons", "LLC", "Group"}

type Company struct {
	namer *Namer
	rand  *rand.Rand
}

func (c *Company) Name() string {
	switch c.rand.Intn(3) {
	case 0:
		return fmt.Sprintf("%s %s", c.namer.LastName(), common.Choose(companySuffixes, c.rand))
	case 1:
		return fmt.Sprint(c.namer.LastName(), "-", c.namer.LastName())
	default:
		return fmt.Sprintf("%s, %s and %s", c.namer.LastName(), c.namer.LastName(),
			c.namer.LastName())
	}
}
