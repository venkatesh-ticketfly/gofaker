package internet

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/gofaker/common"
	"github.com/gofaker/us/company"
	"github.com/gofaker/us/name"
)

var domainSuffixes = []string{"com", "org", "us", "ca", "biz", "info"}

type Email struct {
	company *company.Company
	namer   *name.Namer
	rand    *common.Rand
}

func NewEmail(company *company.Company, namer *name.Namer) *Email {
	return &Email{company, namer, common.Default()}
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
	domainSuffix := e.rand.Choose(domainSuffixes)
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
