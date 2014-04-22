package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

type cardStack struct {
	Title      string      `json:"title"`
	Author     string      `json:"Author"`
	Flashcards []flashcard `json:flashcards`
}

// loadFlashcardStack loads a json file called filename and unmarshals it into the calling struct.
// If there is an error reading the file or unmarshaling it, the error will be returned.
func (stack *cardStack) loadFlashcardStack(filename string) error {
	jsonFileData, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonFileData, stack)

	if err != nil {
		return err
	}
	return nil
}

// Shuffles a cardStack's Flashcards in place using a Fisher-Yates shuffle
func (stack *cardStack) shuffle() {
	rand.Seed(time.Now().UnixNano())

	for i := len(stack.Flashcards) - 1; i > 0; i-- {
		j := rand.Intn(i)
		stack.Flashcards[i], stack.Flashcards[j] = stack.Flashcards[j], stack.Flashcards[i]
	}
}
