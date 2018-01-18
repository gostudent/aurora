package main

import (
	"github.com/gostudent/aurora"
)

func main() {
	quote := []string{"Be yourself", "everyone else is", "already taken."}
	author := "Oscar Wilde"
	filename := "1"
	//Optional argument for gradient color added indexed from 1
	aurora.CreatePort(filename, quote, author, 20)
}
