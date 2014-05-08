package main

import (
	"github.com/alexjohnj/flasher/tbutils"
)

type flashcard struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Image    string `json:"image"`
}

// drawQuestion writes the flashcard's question to the termbox's back buffer.
// It does not clear the display or flush the buffer
func (f *flashcard) drawQuestion() {
	xCoord := tbutils.CalculateXCenterCoord(f.Question)
	tbutils.DrawText(xCoord, 10, f.Question)
}

// drawAnswer writes the flashcard's answer to the termbox's back buffer.
// It does not clear the display or flush the buffer.
func (f *flashcard) drawAnswer() {
	xCoord := tbutils.CalculateXCenterCoord(f.Answer)
	tbutils.DrawText(xCoord, 11, f.Answer)
}
