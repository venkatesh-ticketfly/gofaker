package us

import (
	"fmt"

	"github.com/gofaker/common"
)

var namePrefixes = []string{"Mr.", "Mrs.", "Ms.", "Miss", "Dr."}

var nameSuffixes = []string{"Jr.", "Sr.", "I", "II", "III", "IV", "V", "MD", "DDS", "PhD", "DVM"}

type Namer struct {
	rand *common.Rand
}

func (u *Namer) Locale() string {
	return "US"
}

func (u *Namer) FirstName() string {
	return u.rand.Choose(firstNames)
}

func (u *Namer) LastName() string {
	return u.rand.Choose(lastNames)
}

func (u *Namer) Prefix() string {
	return u.rand.Choose(namePrefixes)
}

func (u *Namer) Suffix() string {
	return u.rand.Choose(nameSuffixes)
}

func (u *Namer) Name() string {
	switch u.rand.Intn(10) {
	case 0:
		return fmt.Sprint(u.Prefix(), " ", u.FirstName(), " ", u.LastName())
	case 1:
		return fmt.Sprint(u.FirstName(), " ", u.LastName(), " ", u.Suffix())
	default:
		return fmt.Sprint(u.FirstName(), " ", u.LastName())
	}
}
