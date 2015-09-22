package us

import (
	"fmt"
	"github.com/gofaker/common"
	"github.com/gofaker/name"
)

const (
	numerifyPattern = '#'
)

var zipFormats = []string{"#####", "#####-####"}
var address2Prefixes = []string{"Apt. ###", "Suite ###"}

type Address struct {
	usNamer name.UsNamer
	rand    *common.Rand
}

func NewAddress(UsNamer name.UsNamer) Address {
	return Address{UsNamer, common.Default()}
}

func (addr *Address) Address1() string {
	streetName := ""
	switch addr.rand.Intn(2) {
	case 0:
		streetName = addr.usNamer.LastName()
	default:
		streetName = addr.usNamer.FirstName()
	}

	streetNum := common.Numerify("####", numerifyPattern, addr.rand)
	return fmt.Sprint(streetNum, " ", streetName, " ", addr.rand.Choose(streetSuffixes))
}

func (addr *Address) Address2() string {
	switch addr.rand.Intn(3) {
	case 0:
		return ""
	default:
		return common.Numerify(addr.rand.Choose(address2Prefixes), numerifyPattern, addr.rand)
	}
}

func (addr *Address) City() string {
	switch addr.rand.Intn(4) {
	case 0:
		return fmt.Sprintf("%s %s%s", addr.cityPrefix(), addr.usNamer.FirstName(),
			addr.citySuffix())
	case 1:
		return fmt.Sprintf("%s %s", addr.cityPrefix(), addr.usNamer.FirstName())
	case 2:
		return fmt.Sprint(addr.usNamer.FirstName(), addr.citySuffix())
	default:
		return fmt.Sprint(addr.usNamer.LastName(), addr.citySuffix())
	}
}

func (addr *Address) State() string {
	return addr.rand.Choose(states)
}

func (addr *Address) StateAbbr() string {
	return addr.rand.Choose(stateAbbrs)
}

func (addr *Address) ZipCode() string {
	return common.Numerify(addr.rand.Choose(zipFormats), numerifyPattern, addr.rand)
}

func (addr *Address) cityPrefix() string {
	return addr.rand.Choose(cityPrefixes)
}

func (addr *Address) citySuffix() string {
	return addr.rand.Choose(citySuffixes)
}
