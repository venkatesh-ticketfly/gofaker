package gofaker

import (
	"github.com/gofaker/us/address"
	"github.com/gofaker/us/company"
	"github.com/gofaker/us/internet"
	"github.com/gofaker/us/name"
)

var UsNamer *name.Namer = name.NewNamer()

var UsAddress *address.Address = address.NewAddress(UsNamer)

var UsCompany *company.Company = company.NewCompany(UsNamer)

var UsEmail *internet.Email = internet.NewEmail(UsCompany, UsNamer)
