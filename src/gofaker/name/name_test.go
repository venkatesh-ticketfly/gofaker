package name

import (
	"gofaker/common/test"
	"testing"
	"fmt"
)

func TestPrefixFirstNameLastNameWhenRandom0(t *testing.T) {
	namer := UsNamer{test.ConstantRand(0)}

	expectedName := fmt.Sprintf("%s %s %s", prefixes[0], firstNames[0], lastNames[0])
	if name := namer.Name(); name != expectedName {
		t.Fatal(name, "should have been", expectedName)
	}
}

func TestFirstNameLastNameSuffixWhenRandom1(t *testing.T) {
	namer := UsNamer{test.ConstantRand(1)}

	expectedName := fmt.Sprintf("%s %s %s", firstNames[1], lastNames[1], suffixes[1])
	if name := namer.Name(); name != expectedName {
		t.Fatal(name, "should have been", expectedName)
	}
}

func TestFirstNameLastNameWhenRandom2(t *testing.T) {
	namer := UsNamer{test.ConstantRand(2)}

	expectedName := fmt.Sprintf("%s %s", firstNames[2], lastNames[2])
	if name := namer.Name(); name != expectedName {
		t.Fatal(name, "should have been", expectedName)
	}
}
