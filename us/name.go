package us

import (
	"fmt"
	"math/rand"

	"github.com/venkatesh-ticketfly/gofaker/common"
)

var namePrefixes = []string{"Mr.", "Mrs.", "Ms.", "Miss", "Dr."}

var nameSuffixes = []string{"Jr.", "Sr.", "I", "II", "III", "IV", "V", "MD", "DDS", "PhD", "DVM"}

type Namer struct {
	rand *rand.Rand
}

func (n *Namer) FirstName() string {
	return common.Choose(firstNames, n.rand)
}

func (n *Namer) LastName() string {
	return common.Choose(lastNames, n.rand)
}

func (n *Namer) Prefix() string {
	return common.Choose(namePrefixes, n.rand)
}

func (n *Namer) Suffix() string {
	return common.Choose(nameSuffixes, n.rand)
}

func (n *Namer) Name() string {
	switch n.rand.Intn(10) {
	case 0:
		return fmt.Sprint(n.Prefix(), " ", n.FirstName(), " ", n.LastName())
	case 1:
		return fmt.Sprint(n.FirstName(), " ", n.LastName(), " ", n.Suffix())
	default:
		return fmt.Sprint(n.FirstName(), " ", n.LastName())
	}
}
