package main

import (
	"bufio"
	"fmt"
	"os"
)

type flashcard struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func (f *flashcard) showCard() {
	fmt.Printf("Q: %s\n(Press the return key to show the answer)\n", f.Question)
	stdinReader := bufio.NewReader(os.Stdin)
	stdinReader.ReadByte()
	fmt.Printf("A: %s\n\n", f.Answer)
}
