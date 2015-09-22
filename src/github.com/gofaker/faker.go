package gofaker

import (
	"github.com/gofaker/address/us"
	"github.com/gofaker/name"
)

var UsNamer name.UsNamer = name.NewUsNamer()

var UsAddress us.Address = us.NewAddress(UsNamer)
