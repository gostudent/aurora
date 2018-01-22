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
