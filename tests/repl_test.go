package tests

import (
	"pokedexcli/internal/cli"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cli.CleanInput(c.input)
		for i := range c.expected {
			want := c.expected[i]
			got := actual[i]

			if want != got {
				t.Errorf("want %s, got %s", want, got)
			}
		}

	}
}
