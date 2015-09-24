package common

import (
	"strconv"
	"strings"
	"testing"
)

func TestNumerifyAllPositionsInPattern(t *testing.T) {
	pattern := "####"
	result := Numerify(pattern, '#', Default())

	if len(pattern) != len(result) {
		t.Fatal("Unmatching length.", "Pattern = ", pattern, " Result = ", result)
	}

	_, err := strconv.Atoi(result)

	if err != nil {
		t.Fatal("Failed to numerify pattern ", pattern, ". Instead returned ", result)
	}
}

func TestNumerifyPatternWithNonReplacebleRune(t *testing.T) {
	pattern := "#####-####"
	result := Numerify(pattern, '#', Default())

	if len(pattern) != len(result) {
		t.Fatal("Unmatching length.", "Pattern = ", pattern, " Result = ", result)
	}

	for _, part_result := range strings.Split(result, "-") {
		_, err := strconv.Atoi(part_result)
		if err != nil {
			t.Fatal("Failed to numerify pattern ", pattern, ". Instead returned ", result)
		}
	}
}

func BenchmarkStringsReplace(b *testing.B) {
	str := "Hello World My World"
	res := strings.Replace(str, " ", "", -1)
	if res != "HelloWorldMyWorld" {
		b.Fatal(res)
	}
}

func BenchmarkStringsMap(b *testing.B) {
	str := "Hello World My World"
	res := strings.Map(func(r rune) rune {
		switch r {
		case rune(32):
			return rune(-1)
		default:
			return r
		}
	},
		str)
	if res != "HelloWorldMyWorld" {
		b.Fatal(res)
	}
}
