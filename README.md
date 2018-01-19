# Aurora

[![OpenCode](https://img.shields.io/badge/Open-Code-ff6a00.svg?style=flat-square)](https://opencode18.github.io)
[![Codacy branch grade](https://img.shields.io/codacy/grade/3cebbd0051d04bce90f0d15f21a092b8/master.svg?style=flat-square)](https://www.codacy.com/app/gostudent/aurora)
[![Aurora](https://img.shields.io/badge/docs-GoDoc-ff69b4.svg?style=flat-square)](https://godoc.org/github.com/gostudent/aurora)
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
