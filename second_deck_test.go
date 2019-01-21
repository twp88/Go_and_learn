package main

import "testing"

func TestDeckPresence(t *testing.T) {
	actual := deck{}

	if actual == nil {
		t.Error("This has not worked")
	}
}
