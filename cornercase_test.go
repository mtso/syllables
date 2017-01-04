package syllables

import (
	"testing"
)

func Test_CornerCases(t *testing.T) {

	cases := []struct {
		want int
		in   string
	}{
		{4, "abalone"},
	}

	for _, c := range cases {
		got, ok := cornercases[c.in]
		if !ok {
			t.Errorf("%q cornercase not found", c.in)
		} else if got != c.want {
			t.Errorf("syllables.In(%q) == %v, expected %v", c.in, got, c.want)
		}
	}
}
