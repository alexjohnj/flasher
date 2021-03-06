# flasher

flasher is a simple flashcard program wrapped up in an easy to use CLI. A deck of flashcards is stored inside of a [TOML][toml-link] file. Flasher reads the TOML file and displays a question and then an answer when the user presses the return key.

flasher is opinionated software. I'm writing the program for when I need to revise so its features will be tailored to my needs.

[toml-link]: https://github.com/toml-lang/toml

## Installation

Assuming you have the [Go development][golang] tools installed and your `$GOPATH` set up, installation is as simple as running these commands:

```bash
go get github.com/alexjohnj/flasher
go install github.com/alexjohnj/flasher
```

[golang]: http://golang.org

## Usage

A deck of flashcards is represented by a TOML file. At the moment you'll need to create the deck manually with a text editor but I plan on adding a simple deck creation tool to flasher in the ~~near~~ future. 

The syntax for the TOML file is as follows:

```toml
title = "The Bridge of Death Test"
author = "Alex Jackson"

[[flashcards]]
question = "What is your name?"
answer = "Arthur, King of the Britons."

[[flashcards]]
question = "What is your quest?"
answer = "To seek the Holy Grail."

[[flashcards]]
question = "What is the airspeed velocity of an unladen swallow?"
answer = "What do you mean, an African or European swallow?"
```

The `author` key is optional but everything else is needed for to load the TOML file. To then test your knowledge, you'd run the following command:

```bash
flasher flash bridge-questions.toml
```

Replacing `bridge-questions.toml` with the name/path to your TOML flashcard deck.

Flasher will then shuffle the deck and display the first question. Press enter to show the answer and enter again to move on to the next question. Press backspace to move back in the deck.

Flasher doesn't try to check if you got an answer right. Since I (and many other people) don't use flashcards with simple one word answers, this would be crazy hard to implement in a user friendly way.

## TODO

There's a few things I want add to flasher, namely:

- Add a flashcard creation command.
- Add the ability to attach an image to a flashcard. This'd likely be implemented by specifying a path/URL to an image in the TOML file and then loading that image with the system's default image previewing application.
- Add some basic stat tracking functionality. Just track which questions you get right/wrong (using an honour system) and add an option to just ask incorrect ones.

