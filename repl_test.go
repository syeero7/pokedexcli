package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Hello  World",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		actualLen := len(actual)
		expectedLen := len(c.expected)

		if actualLen != expectedLen {
			t.Errorf("test failed\nexpected length: %d\nactual length: %d", expectedLen, actualLen)
			return
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := c.expected[i]

			if actualWord != expectedWord {
				t.Errorf("test failed\nexpected word: %s\nactual word: %s", expectedWord, actualWord)
				return
			}

		}
	}

}
