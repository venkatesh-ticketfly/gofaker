package gofaker

import (
	"gofaker/address/us"
	"gofaker/name"
)

var UsNamer name.UsNamer = name.NewUsNamer()

var UsAddress us.Address = us.NewAddress(UsNamer)
