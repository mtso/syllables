package syllables

import (
	"testing"
)

func Test_CountIn(t *testing.T) {
	cases := []struct {
		want int
		in   string
	}{
		// Original test cases
		{3, "syllable"},
		{3, "unicorn"},
		{1, "hi"},
		{2, "hihi"},
		{1, "mmmmmmmmmmmmmmmm"},
		{2, "hoopty"},

		// Additional test cases
		{6, "syllables in this phrase"},
		{31, "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."},
		{102, "If you have just untarred a binary Go distribution, you need to set the environment variable $GOROOT to the full path of the go directory (the one containing this file). You can omit the variable if you unpack it into /usr/local/go, or if you rebuild from sources by running all.bash (see doc/install-source.html). You should also add the Go binary directory $GOROOT/bin to your shell's path."},
		{2, "template"},
		{4, "abalone"},
		{5, "aborigine"},
		{3, "simile"},
		{4, "facsimile"},
		{3, "syncope"},
		{0, ""},
		{1, "yo"},
		{3, "kilogram"},
	}

	for _, c := range cases {
		got := In(c.in)
		if got != c.want {
			t.Errorf("syllables.In(%q) == %v, expected %v", c.in, got, c.want)
		}
	}
}
