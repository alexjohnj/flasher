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
		},
	}
	app.Run(os.Args)
}

func cliFlash(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatal("Incorrect usage. Too few/many arguments") // Useless message, will update
	}

	inputFlashcardStack := new(cardStack)
	err := inputFlashcardStack.loadFlashcardStack(c.Args().First())

	if err != nil {
		log.Fatal(err)
	}

	// Test that everything was read in correctly
	for _, card := range inputFlashcardStack.Flashcards {
		card.showCard()
	}
}
