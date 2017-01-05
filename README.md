# syllables [![Build Status](https://travis-ci.org/mtso/syllables.svg?branch=master)](https://travis-ci.org/mtso/syllables) [![codecov](https://codecov.io/gh/mtso/syllables/branch/master/graph/badge.svg)](https://codecov.io/gh/mtso/syllables) [![GoDoc](https://godoc.org/github.com/mtso/syllables?status.svg)](https://godoc.org/github.com/mtso/syllables)

Go port of the JavaScript syllable counter at https://github.com/wooorm/syllable

## Install

```
go get "github.com/mtso/syllables"
```

## Example

Example usage of `syllables.In(string) int`:

```go
package main

import (
	"fmt"
	"github.com/mtso/syllables"
)

func main() {
	text := "The quick brown fox jumps over the lazy dog."
	syllableCount := syllables.In(text)
	fmt.Printf("There are %v syllables in %q\n", syllableCount, text)
	// Output: There are 11 syllables in "The quick brown fox jumps over the lazy dog."
}
```
