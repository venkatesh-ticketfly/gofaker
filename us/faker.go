package us

import "../common"

type Us struct {
	Namer   *Namer
	Address *Address
	Company *Company
	Email   *Email
}

func NewUs() *Us {
	rand := common.NewRand()
	namer := &Namer{rand}
	company := &Company{namer, rand}
	return &Us{namer, &Address{namer, rand}, company, &Email{company, namer, rand}}
}
