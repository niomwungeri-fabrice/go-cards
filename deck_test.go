package main

import (
	"testing"
)

func TestNewDec(t *testing.T) {
	expected := 16
	d := newDeck()
	if len(d) != expected {
		t.Errorf("Expected %v but received %v", expected, len(d))
	}

	if d[0] != "Ace of Spaces" {
		t.Errorf("Expected Spaces of Ace but received %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Spaces of Ace but received %v", d[len(d)-1])
	}
}
