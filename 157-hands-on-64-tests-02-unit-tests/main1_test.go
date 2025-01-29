package main

import "testing"

func TestSubtract(t *testing.T) {
	x := subtract(1, 1)
	if x != 0 {
		t.Error("Expected", 0, "Got", x)
	}
}
