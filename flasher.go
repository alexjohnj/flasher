package main

import (
	"github.com/codegangsta/cli"
	"log"
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
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "no-shuffle, n",
					Usage: "Presents flashcards in the order they are written in the source JSON file.",
				},
			},
		},
	}
	app.Run(os.Args)
}

func cliFlash(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Print("Incorrect usage")
		cli.ShowCommandHelp(c, "flash")
		os.Exit(1)
	}

	inputFlashcardStack := new(cardStack)
	err := inputFlashcardStack.loadFlashcardStack(c.Args().First())

	if err != nil {
		log.Fatal(err)
	}

	if !c.Bool("no-shuffle") {
		inputFlashcardStack.shuffle()
	}

	// Test that everything was read in correctly
	for _, card := range inputFlashcardStack.Flashcards {
		card.formatCard()
		card.showCard()
	}
}
