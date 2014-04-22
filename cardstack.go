package main

import (
	"encoding/json"
	"io/ioutil"
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
