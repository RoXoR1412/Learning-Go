package main

import "testing"

func TestAdd(t *testing.T) {
	x := add(1, 1)
	if x != 2 {
		t.Error("Expected", 2, "Got", x)
	}
}
