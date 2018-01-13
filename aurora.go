package aurora

import (
	"github.com/ajstarks/svgo"
	"log"
	"math"
	"os"
)

var (
	svgFile *os.File
	err     error
)

/*
create generates an SVG image.
filename - filename without extension (eg:- "sherlock: for DrSeuss.svg)
quote - array of string to be made into quote (eg:-
				["Don't cry because", "it's over,", "smile because", "it happened."])
author - author name (eg:- "Dr. Seuss")
*/
func Create(filename string, quote []string, author string) {
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
	qSpace := int(math.Min(float64((n-2)*90), 90))
	aSpace := int(math.Max(float64((n-3)*150), 150))
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
	canvas.Textlines(width/2, (height/2)-qSpace, quote, 72, 72, "white", "middle")
	canvas.Text(width/2, (height/2)+aSpace, author, aStyle)
	canvas.End()
	svgFile.Close()
}
