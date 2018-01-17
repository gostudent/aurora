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

//New data type for gradient createdd
type gradient struct {
	color1 svg.Offcolor
	color2 svg.Offcolor
}

//An array for all the gradients created.
var gradients = []gradient{
	{svg.Offcolor{0, "#fcdf8a", 1.0}, svg.Offcolor{100, "#F38381", 1.0}},
	{svg.Offcolor{0, "#f54ea2", 1.0}, svg.Offcolor{100, "#FF7676", 1.0}},
	{svg.Offcolor{0, "#17ead9", 1.0}, svg.Offcolor{100, "#6078EA", 1.0}},
	{svg.Offcolor{0, "#622774", 1.0}, svg.Offcolor{100, "#C53364", 1.0}},
	{svg.Offcolor{0, "#7117ea", 1.0}, svg.Offcolor{100, "#EA6060", 1.0}},
	{svg.Offcolor{0, "#42e695", 1.0}, svg.Offcolor{100, "#3BB2B8", 1.0}},
	{svg.Offcolor{0, "#f02fc2", 1.0}, svg.Offcolor{100, "#6094ea", 1.0}},
	{svg.Offcolor{0, "#65799b", 1.0}, svg.Offcolor{100, "#5e2563", 1.0}},
	{svg.Offcolor{0, "#184e68", 1.0}, svg.Offcolor{100, "#57ca85", 1.0}},
	{svg.Offcolor{0, "#5b247a", 1.0}, svg.Offcolor{100, "#1bcedf", 1.0}},
	{svg.Offcolor{0, "#fad961", 1.0}, svg.Offcolor{100, "#f76b1c", 1.0}},
	{svg.Offcolor{0, "#f2d50f", 1.0}, svg.Offcolor{100, "#da0641", 1.0}},
	{svg.Offcolor{0, "#F5515F", 1.0}, svg.Offcolor{100, "#A1051D", 1.0}},
	{svg.Offcolor{0, "#F36265", 1.0}, svg.Offcolor{100, "#961276", 1.0}},
	{svg.Offcolor{0, "#FF57B9", 1.0}, svg.Offcolor{100, "#A704FD", 1.0}},
	{svg.Offcolor{0, "#C56CD6", 1.0}, svg.Offcolor{100, "#3425AF", 1.0}},
	{svg.Offcolor{0, "#13f1fc", 1.0}, svg.Offcolor{100, "#0470dc", 1.0}},
	{svg.Offcolor{0, "#0FF0B3", 1.0}, svg.Offcolor{100, "#036ED9", 1.0}},
	{svg.Offcolor{0, "#c3ec52", 1.0}, svg.Offcolor{100, "#0ba29d", 1.0}},
	{svg.Offcolor{0, "#b1ea4d", 1.0}, svg.Offcolor{100, "#459522", 1.0}},
	{svg.Offcolor{0, "#DFEC51", 1.0}, svg.Offcolor{100, "#73AA0A", 1.0}},
	{svg.Offcolor{0, "#E3E3E3", 1.0}, svg.Offcolor{100, "#5D6874", 1.0}},
	{svg.Offcolor{0, "#CE9FFC", 1.0}, svg.Offcolor{100, "#7367F0", 1.0}},
}

/*CreateSquare generates square SVG image.
filename - filename without extension (eg:- "sherlock: for DrSeuss.svg)
quote - array of string to be made into quote (eg:-
				["Don't cry because", "it's over,", "smile because", "it happened."])
author - author name (eg:- "Dr. Seuss")
*/
func CreateSquare(filename string, quote []string, author string, lgNo ...int) {
	// Setting default gradient if no gradient selected or else selecting the desired gradient.
	lg := []svg.Offcolor{gradients[0].color1, gradients[0].color2}
	if len(lgNo) > 0 {
		lg = []svg.Offcolor{
			gradients[lgNo[0]-1].color1,
			gradients[lgNo[0]-1].color2,
		}
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
func CreateLand(filename string, quote []string, author string, lgNo ...int) {
	// Setting default gradient if no gradient selected or else selecting the desired gradient.
	lg := []svg.Offcolor{gradients[0].color1, gradients[0].color2}
	if len(lgNo) > 0 {
		lg = []svg.Offcolor{
			gradients[lgNo[0]-1].color1,
			gradients[lgNo[0]-1].color2,
		}
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
func CreatePort(filename string, quote []string, author string, lgNo ...int) {
	// Setting default gradient if no gradient selected or else selecting the desired gradient.
	lg := []svg.Offcolor{gradients[0].color1, gradients[0].color2}
	if len(lgNo) > 0 {
		lg = []svg.Offcolor{
			gradients[lgNo[0]-1].color1,
			gradients[lgNo[0]-1].color2,
		}
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
