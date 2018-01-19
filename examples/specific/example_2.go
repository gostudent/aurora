package main

import (
	"github.com/gostudent/aurora"
)

func main() {
	quote := []string{"It is better to", "be hated for what", "you are than to",
		"loved for what", "what you are not."}
	author := "André Gide"
	filename := "2"
	aurora.Create(filename, quote, author, 1000, 2000, 1)
}
