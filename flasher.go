package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/nsf/termbox-go"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "flasher"
	app.Usage = "A simple flashcard app."
	app.Version = "0.3.1"
	app.Author = "Alex Jackson"
	app.Email = "alex@alexj.org"

	app.Commands = []cli.Command{
		{
			Name:      "flash",
			ShortName: "f",
			Usage:     "flasher flash [flashcard-file.toml]",
			Action:    cliFlash,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "no-shuffle, n",
					Usage: "Presents flashcards in the order they are written in the source TOML file.",
				},
			},
		},
		{
			Name:      "info",
			ShortName: "i",
			Usage:     "flasher info [flashcard-file.toml]",
			Action:    cliInfo,
		},
	}
	app.Run(os.Args)
}

func cliFlash(c *cli.Context) {
	// Load flashcards

	if len(c.Args()) != 1 {
		log.Printf("Incorrect usage\n")
		cli.ShowCommandHelp(c, "flash")
		os.Exit(1)
	}

	flashcardStack := new(cardStack)
	err := flashcardStack.loadFlashcardStack(c.Args()[0])

	if err != nil {
		log.Fatalf("Could not load %s (Reason: %s)\n", c.Args()[0], err.Error())
	}

	if len(flashcardStack.Flashcards) < 1 {
		log.Fatalf("Did not find any flashcards in %s!\n", c.Args()[0])
	}

	if !c.Bool("no-shuffle") {
		flashcardStack.shuffle()
	}

	// Init termbox
	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	drawAll(c, flashcardStack)

	//Main Run loop
mainloop:
	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			switch event.Key {
			case termbox.KeyEsc, termbox.KeyCtrlC:
				break mainloop
			case termbox.KeyEnter, termbox.KeyArrowRight:
				flashcardStack.advanceStack()
			case termbox.KeyBackspace2, termbox.KeyArrowLeft:
				flashcardStack.revertStack()
			}

			switch event.Ch {
			case 'q':
				break mainloop

			case 'l':
				flashcardStack.advanceStack()

			case 'h':
				flashcardStack.revertStack()

			case 'r':
				if flashcardStack.atEndOfStack() {
					flashcardStack.StackIndex = 0
					flashcardStack.ShowAnswer = false
				}

			case 'x':
				if flashcardStack.atEndOfStack() {
					flashcardStack.shuffle()
					flashcardStack.StackIndex = 0
					flashcardStack.ShowAnswer = false
				}
			}

		case termbox.EventResize:
			drawAll(c, flashcardStack)
		}

		drawAll(c, flashcardStack)
	}
}

func cliInfo(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Printf("Incorrect usage\n")
		cli.ShowCommandHelp(c, "info")
		os.Exit(1)
	}

	flashcardStack := new(cardStack)
	err := flashcardStack.loadFlashcardStack(c.Args()[0])

	if err != nil {
		log.Fatalf("%s is an invalid file: %s", c.Args()[0], err.Error())
	}

	fmt.Printf("Deck Name: %s\nAuthor: %s\nNumber of Cards: %d\n", flashcardStack.Title, flashcardStack.Author, len(flashcardStack.Flashcards))
}
