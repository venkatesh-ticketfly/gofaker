package name

import (
	"fmt"

	"github.com/gofaker/common"
)

var prefixes = []string{"Mr.", "Mrs.", "Ms.", "Miss", "Dr."}

var suffixes = []string{"Jr.", "Sr.", "I", "II", "III", "IV", "V", "MD", "DDS", "PhD", "DVM"}

func NewNamer() *Namer {
	return New(common.Default())
}

func New(rand *common.Rand) *Namer {
	return &Namer{rand}
}

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
	return u.rand.Choose(prefixes)
}

func (u *Namer) Suffix() string {
	return u.rand.Choose(suffixes)
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
