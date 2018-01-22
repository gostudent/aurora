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
