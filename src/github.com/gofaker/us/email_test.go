package us

import (
	"testing"
)

func TestDomainWordFromStringWithSpaces(t *testing.T) {
	domainWord := formatDomainWord("First Name")
	if domainWord != "first" {
		t.Fatal("Expected first but was ", domainWord)
	}
}

func TestDomainWordFromStringWithNonLetters(t *testing.T) {
	domainWord := formatDomainWord("First5 Name")
	if domainWord != "first" {
		t.Fatal("Expected first but was ", domainWord)
	}
}

func TestDomainWordFromStringWithoutSpaces(t *testing.T) {
	domainWord := formatDomainWord("First-Name")
	if domainWord != "firstname" {
		t.Fatal("Expected first but was ", domainWord)
	}
}
