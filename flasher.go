package main

import (
	"github.com/codegangsta/cli"
	"github.com/nsf/termbox-go"
	"math"
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
	// Init termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	draw()

	// Run loop
mainloop:
	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			switch event.Key {
			case termbox.KeyEsc:
				break mainloop
			}
		case termbox.EventResize:
			draw()
		}
		draw()
	}
}

func draw() {
	w, h := termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Draw termbox border
	termbox.SetCell(0, 0, '+', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(w-1, 0, '+', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(0, h-1, '+', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(w-1, h-1, '+', termbox.ColorDefault, termbox.ColorDefault)

	for x := 1; x < w-1; x++ {
		termbox.SetCell(x, 0, '-', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x, h-1, '-', termbox.ColorDefault, termbox.ColorDefault)
	}

	for y := 1; y < h-1; y++ {
		termbox.SetCell(0, y, '|', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(w-1, y, '|', termbox.ColorDefault, termbox.ColorDefault)
	}

	// Draw a string in the ~centre of the termbox
	message := "Flasher will flash soon!"
	xStart := (w / 2) - int(math.Floor(float64(len(message)/2)))
	for index, runeVal := range message {
		termbox.SetCell(xStart+index, h/2, runeVal, termbox.ColorGreen, termbox.ColorDefault)
	}

	termbox.Flush()
}
