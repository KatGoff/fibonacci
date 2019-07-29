package main

import (
	"testing"
)

func TestFibFinder(t *testing.T) {
	for _, c := range testCases {
		n := FibFinder(c.input)
		if n != c.expected {
			t.Fatalf("FAIL: %s\nCase %q, expected %v, got %v", c.description, c.input, c.expected, n)
		}

		t.Logf("PASS: %v", c.input)
	}
}

var testCases = []struct {
	description string
	input       uint
	expected    uint
}{
	{
		description: "zero",
		input:       0,
		expected:    0,
	},
	{
		description: "one",
		input:       1,
		expected:    1,
	},
	{
		description: "single digit",
		input:       6,
		expected:    8,
	},
	{
		description: "double digit",
		input:       65,
		expected:    17167680177565,
	},
}
