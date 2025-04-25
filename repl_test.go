package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello   world  ",
			expected: []string{"hello", "world"},
		}, {
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}
	testsPassed := 0
	testsFailed := 0
	testsRan := 0

	for _, c := range cases {
		fmt.Println("Running test")
		fmt.Println("====================")
		fmt.Printf("Input string: %s\n", c.input)
		fmt.Printf("Expected output: %v\n\n", c.expected)
		testsRan += 1

		actual := cleanInput(c.input)
		if len(actual) == 0 {
			t.Error("The function returned nothing")
		}

		for i := range c.expected {

			word := actual[i]
			if word == "" {
				continue
			}

			expectedWord := c.expected[i]
			fmt.Printf("expected word: %s\n", expectedWord)
			fmt.Printf("actual word: %s\n", word)
			if word != expectedWord {
				t.Errorf("Word mismatch at position %d for input %q: got %q, expected %q", i, c.input, word, expectedWord)
				testsFailed += 1
			}
		}
		testsPassed += 1
	}
	fmt.Printf("Ran %d tests. Passed %d; Failed %d\n", testsRan, testsPassed, testsFailed)

}
