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
