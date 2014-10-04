package main

import (
	"fmt"
	"github.com/alexjohnj/flasher/tbutils"
	"github.com/codegangsta/cli"
	"github.com/nsf/termbox-go"
)

func drawAll(c *cli.Context, stack *cardStack) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawTermboxChrome(c, stack)
	drawCurrentCard(stack)

	if stack.atEndOfStack() {
		drawEndOfStack()
	}

	termbox.Flush()
}

func drawTermboxChrome(c *cli.Context, stack *cardStack) {
	w, h := termbox.Size()

	// Draw the border
	for x := 0; x <= w; x++ {
		termbox.SetCell(0+x, 0, ' ', termbox.ColorBlack, termbox.ColorWhite)
		termbox.SetCell(0+x, h-1, ' ', termbox.ColorBlack, termbox.ColorWhite)
	}

	// Draw the app name & version
	appString := fmt.Sprintf("%s %s", c.App.Name, c.App.Version)
	tbutils.DrawRichText(2, 0, appString, termbox.ColorBlack, termbox.ColorWhite)

	// Draw the Stack's title
	titleXCoord := tbutils.CalculateXCenterCoord(stack.Title)
	tbutils.DrawRichText(titleXCoord, 0, stack.Title, termbox.ColorBlack, termbox.ColorWhite)

	// Draw the current position in the stack
	fractionalPosition := (float64(stack.StackIndex+1) / float64(len(stack.Flashcards))) * 100
	indexStr := fmt.Sprintf("Card %d of %d (%.0f%%)", stack.StackIndex+1, len(stack.Flashcards), fractionalPosition)
	indexXCoord := tbutils.CalculateXCenterCoord(indexStr)
	tbutils.DrawRichText(indexXCoord, h-1, indexStr, termbox.ColorBlack, termbox.ColorWhite)
}

func drawCurrentCard(stack *cardStack) {
	// Draw the current card
	w, h := termbox.Size()
	currentQuestion := stack.getCurrentFlashcard()

	currentQuestion.drawQuestion()

	if stack.ShowAnswer {
		currentQuestion.drawAnswer()
		// Draw the Q/A divider
		for x := 0; x < w; x++ {
			termbox.SetCell(x, (3 * h / 8), '-', termbox.ColorBlue, termbox.ColorDefault)
		}
	}
}

func drawEndOfStack() {
	_, h := termbox.Size()

	endOfStackMessageLine1 := "End of Stack..."
	endOfStackMessageLine2 := "(q)Quit (r)Restart (x)Reshuffle & Restart."
	eosXCoord1, eosYCoord1 := tbutils.CalculateXCenterCoord(endOfStackMessageLine1), 3*(h/4)
	eosXCoord2, eosYCoord2 := tbutils.CalculateXCenterCoord(endOfStackMessageLine2), (3*(h/4))+1
	tbutils.DrawText(eosXCoord1, eosYCoord1, endOfStackMessageLine1)
	tbutils.DrawText(eosXCoord2, eosYCoord2, endOfStackMessageLine2)
}
