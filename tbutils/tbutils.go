// Package tbutils provides a small number of helper functions for displaying text in a termbox
package tbutils

import (
	"github.com/nsf/termbox-go"
	"math"
	"strings"
)

// DrawText take a string and adds it to the termbox's back buffer.
// If the string contains any newline characters, the string will be split across multiple rows.
// If the string is too long to fit in the termbox, it'll be split across multiple lines to fit.
// This function is the equivalent of calling DrawRichText(x,y,text,termbox.ColorDefault,termbox.ColorDefault)
func DrawText(x, y int, text string) {
	DrawRichText(x, y, text, termbox.ColorDefault, termbox.ColorDefault)
}

// DrawRichText takes a string and adds it to the termbox's back buffer.
// If the string contains any newline characters, it'll be split across multiple rows.
// If the string is too long to find in the termbox, it'll be split across multiple lines to fit.
func DrawRichText(x, y int, text string, fg, bg termbox.Attribute) {
	for row, line := range strings.Split(text, "\n") {
		DrawRichLine(x, y+row, line, fg, bg)
	}
}

// DrawLine adds a single line of text to the termbox's back buffer.
// If the line is too long to fit in the termbox, it will be split across multiple lines.
// This function is the equivalent of calling DrawRichLine(x,y,text,termbox.ColorDefault,termbox.ColorDefault)
func DrawLine(x, y int, text string) {
	DrawRichLine(x, y, text, termbox.ColorDefault, termbox.ColorDefault)
}

// DrawRichLine adds a single line of text to the termbox's back buffer.
// If the line is too long to fit in the termbox, it will be split across multiple lines.
func DrawRichLine(x, y int, text string, fg, bg termbox.Attribute) {
	w, _ := termbox.Size()

	if len(text) > (w - x) {
		DrawRichLine(x, y, text[0:(w-x)], fg, bg)
		DrawRichLine(x, y+1, text[(w-x):len(text)], fg, bg)
	} else {
		for index, ch := range text {
			termbox.SetCell(x+index, y, ch, fg, bg)
		}
	}
}

// CalculateXCenterCoord calculates the x coordinate to write a string to so it is centred in the termbox.
// If the text is longer than the width of the termbox then this function returns 0
func CalculateXCenterCoord(text string) int {
	w, _ := termbox.Size()
	if len(text) > w {
		return 0
	}
	return (w / 2) - int(math.Floor(float64(len(text)/2)))
}
