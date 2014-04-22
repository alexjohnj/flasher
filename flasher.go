package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "flasher"
	app.Version = "0.1.0"
	app.Author = "Alex Jackson"
	app.Email = "alex@alexj.org"

	app.Commands = []cli.Command{
		{
			Name:      "flash",
			ShortName: "f",
			Usage:     "flasher flash [flashcard-file.json]",
			Action:    cliFlash,
		},
	}
	app.Run(os.Args)
}

func cliFlash(c *cli.Context) {
}
