# Aurora

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
	quote := []string{"Be yourself", "everyone else is", "already taken."}
	author := "Oscar Wilde"
	filename := "1"
	aurora.CreatePort(filename, quote, author)
}
```

<img src="examples/2.svg" alt="Smiley face" height="200" width="200">
