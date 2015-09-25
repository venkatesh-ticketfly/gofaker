package us

import "github.com/gofaker/common"

type Us struct {
	Namer *Namer
	Address *Address
	Company *Company
	Email *Email
}

func NewUs() *Us {
	namer := &Namer{common.Default()}
	company :=  &Company{namer, common.Default()}
	return &Us{namer, &Address{namer, common.Default()}, company, &Email{company, namer, common.Default()}}
}
