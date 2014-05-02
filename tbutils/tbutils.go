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
func DrawText(x, y int, text string) {
	for row, line := range strings.Split(text, "\n") {
		DrawLine(x, y+row, line)
	}
}

// DrawLine adds a single line of text to the termbox's back buffer.
// If the line is too long to fit in the termbox, it will be split across multiple lines.
func DrawLine(x, y int, text string) {
	w, _ := termbox.Size()

	if len(text) > (w - x) {
		DrawLine(x, y, text[0:(w-x)])
		DrawLine(x, y+1, text[(w-x):len(text)-1])
	} else {
		for index, char := range text {
			termbox.SetCell(x+index, y, char, termbox.ColorDefault, termbox.ColorDefault)
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
