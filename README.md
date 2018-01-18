# Aurora

![Aurora](https://img.shields.io/badge/docs-GoDoc-ff69b4.svg?style=flat-square)
[![Travis Branch](https://img.shields.io/travis/gostudent/aurora/master.svg?style=flat-square)](https://travis-ci.org/gostudent.aurora)

A Package to create quote images

## Installation

`go get github.com/gostudent/aurora`

## Usage

```go
package main

import (
	"github.com/gostudent/aurora"
)

func main() {
	quote := []string{"It is better to", "be hated for what", "you are than to",
		"loved for what", "you are not."}
	author := "André Gide"
	filename := "something"
	aurora.Create(filename, quote, author)
}
```

<img src="examples/2.svg" alt="Smiley face" height="200" width="200">
