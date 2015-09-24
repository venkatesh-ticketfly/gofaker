package company

import (
	"fmt"

	"github.com/gofaker/common"
	"github.com/gofaker/us/name"
)

var suffixes = []string{"Inc", "and Sons", "LLC", "Group"}

type Company struct {
	namer *name.Namer
	rand  *common.Rand
}

func NewCompany(namer *name.Namer) *Company {
	return &Company{namer, common.Default()}
}

func (c *Company) Name() string {
	switch c.rand.Intn(3) {
	case 0:
		return fmt.Sprintf("%s %s", c.namer.LastName(), c.rand.Choose(suffixes))
	case 1:
		return fmt.Sprint(c.namer.LastName(), "-", c.namer.LastName())
	default:
		return fmt.Sprintf("%s, %s and %s", c.namer.LastName(), c.namer.LastName(),
			c.namer.LastName())
	}
}
