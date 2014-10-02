package main

import (
	"github.com/BurntSushi/toml"
	"math/rand"
	"time"
)

type cardStack struct {
	Title      string
	Author     string
	Flashcards []flashcard
	ShowAnswer bool
	StackIndex int
}

// advanceStack moves forward through the stack of flashcards.
// If the current flashcard isn't showing an answer, it will flag it to show the answer.
// Otherwise it will increment the StackIndex and disable the ShownAnswer flag.
// This function will do nothing if incrementing the StackIndex would produce an OOB error.
func (s *cardStack) advanceStack() {
	if s.StackIndex >= len(s.Flashcards)-1 && s.ShowAnswer {
		return
	}

	if s.ShowAnswer {
		s.StackIndex++
		s.ShowAnswer = false
	} else if !s.ShowAnswer {
		s.ShowAnswer = true
	}
}

// revertStack is the inverse of advanceStack
// This function will do nothing if decrementing StackIndex would produce an OOB error.
func (s *cardStack) revertStack() {
	if s.ShowAnswer {
		s.ShowAnswer = false
	} else if !s.ShowAnswer && s.StackIndex > 0 {
		s.StackIndex--
		s.ShowAnswer = true
	}
}

// getCurrentFlashcard is a convenience function that returns a copy of the current flashcard.
func (s *cardStack) getCurrentFlashcard() flashcard {
	return s.Flashcards[s.StackIndex]
}

// loadFlashcardStack loads a json file called filename and unmarshals it into the calling struct.
// If there is an error reading the file or unmarshaling it, the error will be returned.
func (stack *cardStack) loadFlashcardStack(filename string) error {
	_, err := toml.DecodeFile(filename, &stack)

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
