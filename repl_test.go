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
			input:    "  Hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Hello  world world world",
			expected: []string{"hello", "world", "world", "world"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("Slice lengths do not match!")	
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words dont match!")	
			}
		}
	}
}

