package main

import (
	"encoding/csv"
	"fmt"
	"github.com/codegangsta/cli"
	"io"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "psv"
	app.Usage = "parse csv"

	app.Action = func(c *cli.Context) {
		var fp *os.File
		la := len(c.Args())

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
			fmt.Println(record)
		}
	}

	app.Run(os.Args)
}
