package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type flashcard struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Image    string `json:"image"`
}

func (f *flashcard) showCard() {
	fmt.Printf("Q: %s\n(Press the return key to show the answer)\n", f.Question)
	stdinReader := bufio.NewReader(os.Stdin)
	stdinReader.ReadByte()
	fmt.Printf("A: %s\n\n", f.Answer)
}

// formatCard replaces a small subset of Markdown sequences with ANSI escape codes.
// ** and __ get replaced with the ANSI bold sequence.
func (f *flashcard) formatCard() {
	// Match **$1**
	re := regexp.MustCompile("\\*\\*(.+?)\\*\\*")
	f.Question = re.ReplaceAllString(f.Question, "\033[1m$1\033[0m")
	f.Answer = re.ReplaceAllString(f.Answer, "\033[1m$1\033[0m")

	// Match **__$1__**
	re = regexp.MustCompile("__(.+?)__")
	f.Question = re.ReplaceAllString(f.Question, "\033[1m$1\033[0m")
	f.Answer = re.ReplaceAllString(f.Answer, "\033[1m$1\033[0m")
}
