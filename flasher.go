package main

import (
	"fmt"
	"github.com/alexjohnj/flasher/tbutils"
	"github.com/codegangsta/cli"
	"github.com/nsf/termbox-go"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "flasher"
	app.Version = "0.3.0"
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
		{
			Name:      "info",
			ShortName: "i",
			Usage:     "flasher info [flashcard-file.json]",
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
		log.Fatal(err)
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
	drawAll(flashcardStack)

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
				drawAll(flashcardStack)
			case termbox.KeyBackspace2, termbox.KeyArrowLeft:
				flashcardStack.revertStack()
				drawAll(flashcardStack)
			}

			switch event.Ch {
			case 'q':
				break mainloop

			case 'l':
				flashcardStack.advanceStack()
				drawAll(flashcardStack)

			case 'h':
				flashcardStack.revertStack()
				drawAll(flashcardStack)

			case 'r':
				if flashcardStack.atEndOfStack() {
					flashcardStack.StackIndex = 0
					flashcardStack.ShowAnswer = false
					drawAll(flashcardStack)
				}

			case 'x':
				if flashcardStack.atEndOfStack() {
					flashcardStack.shuffle()
					flashcardStack.StackIndex = 0
					flashcardStack.ShowAnswer = false
					drawAll(flashcardStack)
				}
			}

		case termbox.EventResize:
			drawAll(flashcardStack)
		}

		drawAll(flashcardStack)
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

func drawAll(stack *cardStack) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	w, h := termbox.Size()

	// Draw termbox border
	termbox.SetCell(0, 0, '+', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(w-1, 0, '+', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(0, h-1, '+', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(w-1, h-1, '+', termbox.ColorDefault, termbox.ColorDefault)

	for x := 1; x < w-1; x++ {
		termbox.SetCell(x, 0, '-', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x, h-1, '-', termbox.ColorDefault, termbox.ColorDefault)
	}

	// Draw the Stack's title
	titleXCoord := tbutils.CalculateXCenterCoord(stack.Title)
	tbutils.DrawText(titleXCoord, 1, stack.Title)

	// Draw the current card
	currentQuestion := stack.getCurrentFlashcard()
	currentQuestion.drawQuestion()
	if stack.ShowAnswer {
		currentQuestion.drawAnswer()
		// Draw the Q/A divider
		for x := 0; x < w; x++ {
			termbox.SetCell(x, (3 * h / 8), '-', termbox.ColorBlue, termbox.ColorDefault)
		}
	}

	// Draw end of stack message
	if stack.atEndOfStack() {
		endOfStackMessageLine1 := "End of Stack..."
		endOfStackMessageLine2 := "(q)Quit (r)Restart (x)Reshuffle & Restart."
		eosXCoord1, eosYCoord1 := tbutils.CalculateXCenterCoord(endOfStackMessageLine1), (3 * (h / 4))
		eosXCoord2, eosYCoord2 := tbutils.CalculateXCenterCoord(endOfStackMessageLine2), (3*h/4)+1
		tbutils.DrawText(eosXCoord1, eosYCoord1, endOfStackMessageLine1)
		tbutils.DrawText(eosXCoord2, eosYCoord2, endOfStackMessageLine2)
	}

	// Draw the current index
	indexStr := fmt.Sprintf("(%d/%d)", stack.StackIndex+1, len(stack.Flashcards))
	indexXCoord, indexYCoord := tbutils.CalculateXCenterCoord(indexStr), h-1
	tbutils.DrawText(indexXCoord, indexYCoord, indexStr)

	// Write out the back buffer
	termbox.Flush()
}
