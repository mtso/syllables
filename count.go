// Package syllables provides a Go implementation of 
// the syllable counter from `github.com/wooorm/syllable`
package syllables

import (
	"strings"
)

type counter struct {
	count, index, length int
	singular             string
	parts                []string
}

// Returns the integer count of syllables in the input byte array.
func InBytes(b []byte) int {
	s := string(b[:len(b)])
	return In(s)
}

// Returns the integer count of syllables in the input string.
func In(text string) int {

	// Prepare input text by converting to lowercase
	// and removing all non-alphabetic runes
	text = strings.ToLower(text)
	text = expressionNonalphabetic.ReplaceAllString(text, "")

	// Return early when possible
	if len(text) < 1 {
		return 0
	}
	if len(text) < 3 {
		return 1
	}

	// If value is part of cornercases,
	// return hardcoded value
	if syllables, ok := cornercases[text]; ok {
		return syllables
	}

	// Initialize counter
	c := counter{}

	// Count and remove matched prefixes and suffixes
	text = expressionTriple.ReplaceAllStringFunc(text, c.countAndRemove(3))
	text = expressionDouble.ReplaceAllStringFunc(text, c.countAndRemove(2))
	text = expressionSingle.ReplaceAllStringFunc(text, c.countAndRemove(1))

	// Count multiple consanants
	c.parts = consanants.Split(text, -1)
	c.index = 0
	c.length = len(c.parts)

	for ; c.index < c.length; c.index++ {
		if c.parts[c.index] != "" {
			c.count++
		}
	}

	// Subtract one for maches which should be
	// counted as one but are counted as two
	subtractOne := c.countInPlace(-1)
	expressionMonosyllabicOne.ReplaceAllStringFunc(text, subtractOne)
	expressionMonosyllabicTwo.ReplaceAllStringFunc(text, subtractOne)

	// Add one for maches which should be
	// counted as two but are counted as one
	addOne := c.countInPlace(1)
	expressionDoubleSyllabicOne.ReplaceAllStringFunc(text, addOne)
	expressionDoubleSyllabicTwo.ReplaceAllStringFunc(text, addOne)
	expressionDoubleSyllabicThree.ReplaceAllStringFunc(text, addOne)
	expressionDoubleSyllabicFour.ReplaceAllStringFunc(text, addOne)

	if c.count < 1 {
		return 1
	}
	return c.count
}

func (c *counter) countAndRemove(increment int) func(string) string {
	return func(in string) string {
		c.count += increment
		return ""
	}
}

func (c *counter) countInPlace(increment int) func(string) string {
	return func(in string) string {
		c.count += increment
		return in
	}
}
