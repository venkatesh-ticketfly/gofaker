package name

import (
	"fmt"
	"gofaker/common"
)

var prefixes = []string{"Mr.", "Mrs.", "Ms.", "Miss", "Dr."}

var suffixes = []string{"Jr.", "Sr.", "I", "II", "III", "IV", "V", "MD", "DDS", "PhD", "DVM"}

func NewUsNamer() UsNamer {
	return New(common.Default())
}

func New(rand *common.Rand) UsNamer {
	return UsNamer{rand}
}

type UsNamer struct {
	rand *common.Rand
}

func (u *UsNamer) Locale() string {
	return "US"
}

func (u *UsNamer) FirstName() string {
	return u.rand.Choose(firstNames)
}

func (u *UsNamer) LastName() string {
	return u.rand.Choose(lastNames)
}

func (u *UsNamer) Prefix() string {
	return u.rand.Choose(prefixes)
}

func (u *UsNamer) Suffix() string {
	return u.rand.Choose(suffixes)
}

func (u *UsNamer) Name() string {
	switch u.rand.Intn(10) {
	case 0:
		return fmt.Sprint(u.Prefix(), " ", u.FirstName(), " ", u.LastName())
	case 1:
		return fmt.Sprint(u.FirstName(), " ", u.LastName(), " ", u.Suffix())
	default:
		return fmt.Sprint(u.FirstName(), " ", u.LastName())
	}
}
