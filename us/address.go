package us

import (
	"fmt"
	"math/rand"

	"github.com/venkatesh-ticketfly/gofaker/common"
)

const (
	numerifyPattern = '#'
)

var zipFormats = []string{"#####", "#####-####"}
var address2Prefixes = []string{"Apt. ###", "Suite ###"}

type Address struct {
	namer *Namer
	rand  *rand.Rand
}

func (addr *Address) Address1() string {
	streetName := ""
	switch addr.rand.Intn(2) {
	case 0:
		streetName = addr.namer.LastName()
	default:
		streetName = addr.namer.FirstName()
	}

	streetNum := common.Numerify("####", numerifyPattern, addr.rand)
	return fmt.Sprint(streetNum, " ", streetName, " ", common.Choose(streetSuffixes, addr.rand))
}

func (addr *Address) Address2() string {
	switch addr.rand.Intn(3) {
	case 0:
		return ""
	default:
		addr2Prefix := common.Choose(address2Prefixes, addr.rand)
		return common.Numerify(addr2Prefix, numerifyPattern, addr.rand)
	}
}

func (addr *Address) City() string {
	switch addr.rand.Intn(4) {
	case 0:
		return fmt.Sprintf("%s %s%s", addr.cityPrefix(), addr.namer.FirstName(),
			addr.citySuffix())
	case 1:
		return fmt.Sprintf("%s %s", addr.cityPrefix(), addr.namer.FirstName())
	case 2:
		return fmt.Sprint(addr.namer.FirstName(), addr.citySuffix())
	default:
		return fmt.Sprint(addr.namer.LastName(), addr.citySuffix())
	}
}

func (addr *Address) State() string {
	return common.Choose(states, addr.rand)
}

func (addr *Address) StateAbbr() string {
	return common.Choose(stateAbbrs, addr.rand)
}

func (addr *Address) ZipCode() string {
	zipFormat := common.Choose(zipFormats, addr.rand)
	return common.Numerify(zipFormat, numerifyPattern, addr.rand)
}

func (addr *Address) cityPrefix() string {
	return common.Choose(cityPrefixes, addr.rand)
}

func (addr *Address) citySuffix() string {
	return common.Choose(citySuffixes, addr.rand)
}
