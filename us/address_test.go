package us

import (
	"fmt"
	"testing"

	"../common/test"
)

func TestZipCodeReturns5DigitFormat(t *testing.T) {
	addr := Address{&Namer{}, test.ConstantRand(0)}

	zipCode := addr.ZipCode()
	if zipCode != "00000" {
		t.Fatal("Expected to return 5 digit zip code, but returned ", zipCode)
	}
}

func TestZipCodeReturns5DigitHyphen4DigitFormat(t *testing.T) {
	addr := Address{&Namer{}, test.ConstantRand(1)}

	zipCode := addr.ZipCode()
	if zipCode != "11111-1111" {
		t.Fatal("Expected to return 5 digit zip code, but returned ", zipCode)
	}
}

func TestCityNameLongForm(t *testing.T) {
	rand := test.ConstantRand(0)
	namer := &Namer{rand}
	addr := Address{namer, rand}

	city := addr.City()
	expectedCity := fmt.Sprint(cityPrefixes[0], " ", namer.FirstName(), citySuffixes[0])
	if city != expectedCity {
		t.Fatal("Expected ", expectedCity, " but was ", city)
	}
}

func TestCityNameWithPrefixOnly(t *testing.T) {
	rand := test.ConstantRand(1)
	namer := &Namer{rand}
	addr := Address{namer, rand}

	city := addr.City()
	expectedCity := fmt.Sprint(cityPrefixes[1], " ", namer.FirstName())
	if city != expectedCity {
		t.Fatal("Expected ", expectedCity, " but was ", city)
	}
}

func TestCityNameWithSuffixOnly(t *testing.T) {
	rand := test.ConstantRand(2)
	namer := &Namer{rand}
	addr := Address{namer, rand}

	city := addr.City()
	expectedCity := fmt.Sprint(namer.FirstName(), citySuffixes[2])
	if city != expectedCity {
		t.Fatal("Expected ", expectedCity, " but was ", city)
	}
}

func TestCityLastNameWithSuffixOnly(t *testing.T) {
	rand := test.ConstantRand(3)
	namer := &Namer{rand}
	addr := Address{namer, rand}

	city := addr.City()
	expectedCity := fmt.Sprint(namer.LastName(), citySuffixes[3])
	if city != expectedCity {
		t.Fatal("Expected ", expectedCity, " but was ", city)
	}
}
