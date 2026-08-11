package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "Pikachu",
			expected: []string{"pikachu"},
		},
		{
			input:    "  Charmander  Bulbasaur   PIKACHU  ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "charizard\tWEEDLE\r\nmewtwo",
			expected: []string{"charizard", "weedle", "mewtwo"},
		},
		{
			input:    "EXPLORE  cavern-of-rumours\tcook 1 2 3",
			expected: []string{"explore", "cavern-of-rumours", "cook", "1", "2", "3"},
		},
		{
			input:    "hElP",
			expected: []string{"help"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("got %d words, want %d: %v", len(actual), len(c.expected), actual)
			continue
		}
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			// error and continue here
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}
