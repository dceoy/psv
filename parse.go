package main

import (
	"encoding/csv"
	"fmt"
	"github.com/codegangsta/cli"
	"io"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "psv"
	app.Usage = "parse csv"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "separator, s",
			Value: ",",
			Usage: "separater for fields",
		},
	}

	app.Action = func(c *cli.Context) {
		var fp *os.File
		la := len(c.Args())
		sep := c.String("separator")

		switch {
		case la == 0:
			fp = os.Stdin
		case la == 1:
			var err error
			fp, err = os.Open(c.Args()[0])
			if err != nil {
				panic(err)
			}
			defer fp.Close()
		case la >= 2:
			println("psv: too many arguments")
		}

		reader := csv.NewReader(fp)
		reader.Comma = ','
		reader.LazyQuotes = true
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			fmt.Println(strings.Join(record, sep))
		}
	}

	app.Run(os.Args)
}
