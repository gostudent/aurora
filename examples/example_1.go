package main

import (
	"github.com/gostudent/aurora"
)

func main() {
	quote := []string{"Be yourself;", "everyone else is", "is already taken."}
	author := "Oscar Wilde"
	filename := "3"
	aurora.CreatePort(filename, quote, author)
}
