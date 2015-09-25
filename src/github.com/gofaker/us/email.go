package us

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode"

	"github.com/gofaker/common"
)

var domainSuffixes = []string{"com", "org", "us", "ca", "biz", "info"}

type Email struct {
	company *Company
	namer   *Namer
	rand    *rand.Rand
}

func (e *Email) UserName() string {
	switch e.rand.Intn(2) {
	case 0:
		return strings.ToLower(e.namer.FirstName())
	default:
		return strings.ToLower(fmt.Sprint(e.namer.FirstName(), "_", e.namer.LastName()))
	}
}

func (e *Email) Email() string {
	domainWord := formatDomainWord(e.company.Name())
	domainSuffix := common.Choose(domainSuffixes, e.rand)
	return fmt.Sprintf("%s@%s.%s", e.UserName(), domainWord, domainSuffix)
}

func formatDomainWord(s string) string {
	domainWord := strings.Split(s, " ")[0]
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return unicode.ToLower(r)
		} else {
			return rune(-1)
		}
	}, domainWord)
}
