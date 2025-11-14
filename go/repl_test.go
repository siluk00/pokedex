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
		// add more casees here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("got %d words, wanted %d words", len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord { // Here we use the variables
				t.Errorf("got word %q, wanted %q", word, expectedWord)
			}
		}

		for _, c := range cases {
			actual := cleanInput(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("got %d words, wanted %d words", len(actual), len(c.expected))
				continue
			}
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				if word != expectedWord { // Here we use the variables
					t.Errorf("got word %q, wanted %q", word, expectedWord)
				}
			}
		}
	}
}
