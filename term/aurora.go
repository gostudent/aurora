package main

import (
	"os"
	"strings"

	"github.com/gostudent/aurora"
	"gopkg.in/urfave/cli.v1"
)

/*
quoteMaker - takes the argument passed in the terminal under  'q' flag  then creates a slice with four words(max) in each element so there
			 will be max 4 word in one line of the image

			 qwords- is the slice consisting of all words
			 qslice- the resulting slice which has 4 words(max) in each element
*/
func quoteMaker(quote string) []string {
	var qslice []string
	qwords := strings.Split(quote, " ")

	var quotepart string

	for i, str := range qwords {
		quotepart += str
		quotepart += " "

		if i > 0 {
			if i%4 == 0 {
				qslice = append(qslice, quotepart)
				quotepart = ""
			}
		}
	}
	return qslice
}

func main() {

	app := cli.NewApp()

	//Defining the flags
	app.Flags = []cli.Flag{

		cli.StringFlag{
			Name:  "quote,q",
			Value: "Default quote template.",
			Usage: "quote that will go in the image",
		},
		cli.StringFlag{
			Name:  "author,a",
			Value: "Anonymous",
			Usage: "Author of the quote",
		},
		cli.StringFlag{
			Name:  "filename,f",
			Value: "quote",
			Usage: "name of the file",
		},
		cli.StringFlag{
			Name:  "orientation,o",
			Value: "0",
			Usage: " Orientation. 0 - Square (600 X 600) , 1 - Landscape  and  2 - Portrait",
		},
		cli.StringFlag{
			Name:  "gradient,g",
			Value: "0",
			Usage: "Gradient : 0-22 integer input.",
		},
	}

	app.Action = func(c *cli.Context) error {

		quote := []string{"Default quote", "template."}
		author := "Anonymous"
		filename := "quote"
		quoteargument := ""
		orientation := 0
		gradientInt := 0

		if len(c.GlobalFlagNames()) > 0 {

			//taking all the arguments
			quoteargument = c.String("q")
			author = c.String("a")
			filename = c.String("f")
			orientation = c.Int("o")
			gradientInt = c.Int("g")

			//making the quote string argument into a slice of strings with 4 words (max) in each element
			quote = quoteMaker(quoteargument)

			gradient := aurora.Gradients[gradientInt].Colors

			if orientation == 0 {
				aurora.Create(filename, quote, author, 600, 600, gradient) // h=600 and w=600

			} else if orientation == 1 {
				aurora.CreateLand(filename, quote, author, gradient)

			} else if orientation == 2 {
				aurora.CreatePort(filename, quote, author, gradient)
			}

		}

		return nil
	}

	app.Run(os.Args)

}
