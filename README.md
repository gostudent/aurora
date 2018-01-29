# Aurora

[![Go Report Card](https://goreportcard.com/badge/github.com/gostudent/aurora?style=flat-square)](https://goreportcard.com/report/github.com/gostudent/aurora)
[![OpenCode](https://img.shields.io/badge/Open-Code-ff6a00.svg?style=flat-square)](https://opencode18.github.io)
[![Codacy branch grade](https://img.shields.io/codacy/grade/3cebbd0051d04bce90f0d15f21a092b8/master.svg?style=flat-square)](https://www.codacy.com/app/gostudent/aurora)
[![Aurora](https://img.shields.io/badge/docs-GoDoc-ff69b4.svg?style=flat-square)](https://godoc.org/github.com/gostudent/aurora)
[![Travis Branch](https://img.shields.io/travis/gostudent/aurora/master.svg?style=flat-square)](https://travis-ci.org/gostudent.aurora)

A Package to create quote images

## Installation

`go get github.com/gostudent/aurora`

## Usage

### To Generate Square Image by Using a Custom Gradient:

```go
package main

import (
	"github.com/gostudent/aurora"
)

func main() {
	quote := []string{"It is better to", "be hated for what", "you are than to",
		"loved for what", "what you are not."}
	author := "André Gide"
	filename := "2"
	lg := []aurora.Color{{0, "#65799b", 1.0}, {25, "#5e2563", 1.0}, {75, "#66666", 1.0}, {100, "#fafafa", 1.0}}
	aurora.CreateSquare(filename, quote, author, lg)
}
```

<img src="examples/square/2.svg" alt="Smiley face" height="200" width="200">


### To Generate Landscape Image by Using a Pre-defined Gradient

```go
package main

import (
	"github.com/gostudent/aurora"
)

func main() {
	quote := []string{"It is better to", "be hated for what", "you are than to",
		"loved for what", "what you are not."}
	author := "André Gide"
	filename := "2"
	lg := aurora.Gradients[7].Colors
	aurora.CreateLand(filename, quote, author, lg)
}
```

<img src="examples/landscape/2.svg" alt="Smiley face" height="240" width="320">


### To Generate Portrait Image by Using a Pre-defined Gradient

```go
package main

import (
	"github.com/gostudent/aurora"
)

func main() {
	quote := []string{"Be yourself", "everyone else is", "already taken."}
	author := "Oscar Wilde"
	filename := "1"
	lg := aurora.Gradients[7].Colors
	aurora.CreatePort(filename, quote, author, lg)
}
```

<img src="examples/portrait/1.svg" alt="Smiley face" height="320" width="240">

### To Generate Image with Custom Dimensions 

```go
package main

import (
	"github.com/gostudent/aurora"
)

func main() {
	quote := []string{"It is better to", "be hated for what", "you are than to",
		"loved for what", "what you are not."}
	author := "André Gide"
	filename := "2"
	lg := aurora.Gradients[7].Colors
	aurora.Create(filename, quote, author, 1000, 2000, lg)
}
```

<img src="examples/specific/2.svg" alt="Smiley face" height="200" width="400">
