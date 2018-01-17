package aurora

import (
	"log"
	"os"

	"github.com/ajstarks/svgo"
)

var (
	svgFile *os.File
	err     error
)

/*
CreateSquare generates a square SVG image.
filename - filename without extension (eg:- "sherlock: for DrSeuss.svg)
quote - array of string to be made into quote (eg:-
				["Don't cry because", "it's over,", "smile because", "it happened."])
author - author name (eg:- "Dr. Seuss")
*/
func CreateSquare(filename string, quote []string, author string) {
	// Linear Gradient Colors
	lg := []svg.Offcolor{
		{0, "#17ead9", 1.0},
		{100, "#6078ea", 1.0},
	}
	// Adding Extension
	filename += ".svg"
	// Creating File
	svgFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Adding Blockquotes to Quote Text
	quote[0] = "\"" + quote[0]
	quote[len(quote)-1] = quote[len(quote)-1] + "\""
	// Finding Pixel Spacing
	n := len(quote)
	var qSpace int
	//To make quote block vertically centered
	if n%2 != 0 {
		qSpace = int((float64(n) / float64(2))) * 75
	} else {
		qSpace = int((float64(n)/float64(2))-1) * 75
	}
	var aSpace int
	//To make Author appear just below quote block vertically centered
	if n%2 != 0 {
		aSpace = int((float64(n)/float64(2))+1) * 75
	} else {
		aSpace = int((float64(n)/float64(2))+1) * 75
	}
	// aSpace := int(math.Max(float64((n-3)*150), 150))
	aStyle := "text-anchor:middle;font-size:54px;fill:white;font-style:italic;"
	// Height and Width
	width := 1000
	height := 1000
	// Drawing Starts
	canvas := svg.New(svgFile)
	canvas.Start(width, height)
	canvas.Def()
	canvas.LinearGradient("bb", 0, 0, 100, 100, lg)
	canvas.DefEnd()
	canvas.Square(0, 0, 1000, "fill: url(#bb)")
	canvas.Textlines(width/2, (height/2)-qSpace, quote, 72, 75, "white", "middle")
	canvas.Text(width/2, (height/2)+aSpace, author, aStyle)
	canvas.End()
	svgFile.Close()
}

/*
CreateLand generates a landscape SVG image.
filename - filename without extension (eg:- "sherlock: for DrSeuss.svg)
quote - array of string to be made into quote (eg:-
				["Don't cry because", "it's over,", "smile because", "it happened."])
author - author name (eg:- "Dr. Seuss")
*/
func CreateLand(filename string, quote []string, author string) {
	// Linear Gradient Colors
	lg := []svg.Offcolor{
		{0, "#17ead9", 1.0},
		{100, "#6078ea", 1.0},
	}
	// Adding Extension
	filename += ".svg"
	// Creating File
	svgFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Adding Blockquotes to Quote Text
	quote[0] = "\"" + quote[0]
	quote[len(quote)-1] = quote[len(quote)-1] + "\""
	// Finding Pixel Spacing
	n := len(quote)
	//To make quote block vertically centered
	var qSpace int
	if n%2 != 0 {
		qSpace = int((float64(n) / float64(2))) * 75
	} else {
		qSpace = int((float64(n)/float64(2))-1) * 75
	}
	var aSpace int
	//To make Author appear just below quote block vertically centered
	if n%2 != 0 {
		aSpace = int((float64(n)/float64(2))+1) * 75
	} else {
		aSpace = int((float64(n)/float64(2) + 1)) * 75
	}
	aStyle := "text-anchor:middle;font-size:54px;fill:white;font-style:italic;"
	// Height and Width
	width := 1600
	height := 1200
	// Drawing Starts
	canvas := svg.New(svgFile)
	canvas.Start(width, height)
	canvas.Def()
	canvas.LinearGradient("bb", 0, 0, 100, 100, lg)
	canvas.DefEnd()
	canvas.Rect(0, 0, width, height, "fill: url(#bb)")
	canvas.Textlines(width/2, (height/2)-qSpace, quote, 72, 75, "white", "middle")
	canvas.Text(width/2, (height/2)+aSpace, author, aStyle)
	canvas.End()
	svgFile.Close()
}

/*
CreateLand generates a portrait SVG image.
filename - filename without extension (eg:- "sherlock: for DrSeuss.svg)
quote - array of string to be made into quote (eg:-
				["Don't cry because", "it's over,", "smile because", "it happened."])
author - author name (eg:- "Dr. Seuss")
*/
func CreatePort(filename string, quote []string, author string) {
	// Linear Gradient Colors
	lg := []svg.Offcolor{
		{0, "#17ead9", 1.0},
		{100, "#6078ea", 1.0},
	}
	// Adding Extension
	filename += ".svg"
	// Creating File
	svgFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Adding Blockquotes to Quote Text
	quote[0] = "\"" + quote[0]
	quote[len(quote)-1] = quote[len(quote)-1] + "\""
	// Finding Pixel Spacing
	n := len(quote)
	//To make quote block vertically centered
	var qSpace int
	if n%2 != 0 {
		qSpace = int((float64(n) / float64(2))) * 75
	} else {
		qSpace = int((float64(n)/float64(2))-1) * 75
	}
	var aSpace int
	//To make Author appear just below quote block vertically centered
	if n%2 != 0 {
		aSpace = int((float64(n)/float64(2))+1) * 75
	} else {
		aSpace = int((float64(n)/float64(2))+1) * 75
	}
	aStyle := "text-anchor:middle;font-size:54px;fill:white;font-style:italic;"
	// Height and Width
	width := 1200
	height := 1600
	// Drawing Starts
	canvas := svg.New(svgFile)
	canvas.Start(width, height)
	canvas.Def()
	canvas.LinearGradient("bb", 0, 0, 100, 100, lg)
	canvas.DefEnd()
	canvas.Rect(0, 0, width, height, "fill: url(#bb)")
	canvas.Textlines(width/2, (height/2)-qSpace, quote, 72, 75, "white", "middle")
	canvas.Text(width/2, (height/2)+aSpace, author, aStyle)
	canvas.End()
	svgFile.Close()
}
