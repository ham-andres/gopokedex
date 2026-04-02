package main

import ("testing"
	
	)

func TestCleanInput(t *testing.T){
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "",
			expected: []string{},
		},
		{
			input: " hello world  ",
			expected: []string{"hello","world"},
		},
		{
			input: " Hello WorLD",
			expected: []string{"hello", "world"},
		},
		{
			input: "hELLO wORLD",
			expected: []string{"hello","world"},
		},

	}

	for _, c:= range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
		t.Errorf("length doesnt match")	
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual and expected mismatch")
			}
		}
	}

}
