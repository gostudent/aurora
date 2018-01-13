package main

import (
	"github.com/gostudent/aurora"
)

func main() {
	quote := []string{"Be yourself;", "everyone else is", "is already taken."}
	author := "Oscar Wilde"
	filename := "1"
	aurora.Create(filename, quote, author)
}
