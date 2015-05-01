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
	app.Usage = "Parse CSV"
	app.Version = "0.2.0"
	app.Author = "dceoy"
	app.Email = "d.narsil@gmail.com"
	app.HideHelp = true

	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} [options] [arguments...]

VERSION:
   {{.Version}}{{if or .Author .Email}}

AUTHOR:{{if .Author}}
   {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
   {{.Email}}{{end}}{{end}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`

	app.Flags = []cli.Flag{
		cli.HelpFlag,
		cli.StringFlag{
			Name:  "separator, s",
			Value: ",",
			Usage: "output separator for fields",
		},
	}

	app.Action = parse

	app.Run(os.Args)
}

func parse(c *cli.Context) {
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
