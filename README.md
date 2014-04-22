# flasher

flasher is a simple flashcards program wrapped up in an easy to use CLI. A deck of flashcards is stored inside of a JSON file. Flasher reads the JSON file and displays a question and then an answer when the user presses the return key. 

flasher is/will be opinionated software. I'm writing the program for when I need to revise so its features will be tailored to my needs.

## Installation

Assuming you have the [Go development][golang] tools installed and your `$GOPATH` set up, installation is as simple as running these commands:

```bash
go get github.com/alexjohnj/flasher
go install github.com/alexjohnj/flasher
```

[golang]: http://golang.org

## Usage

A deck of flashcards is represented by a JSON file. At the moment you'll need to create the deck manually with a text editor but I plan on adding a simple deck creation tool to flasher in the near future. 

The syntax for the JSON file is as follows:

```json
{
  "title": "The Bridge of Death Test",
  "flashcards": [
    {
        "question": "What is your name?",
        "answer": "Arthur, King of the Britons."
    },
    {
        "question" : "What is your quest?",
        "answer" : "To seek the Holy Grail."
    },
    {
        "question" : "What is the airspeed velocity of an unladen swallow?",
        "answer": "What do you mean, an African or European swallow?"
    },
  ]
}
```

The `title` key is optional but the `flashcards` array needs to be implemented because, well, that's the whole point of flasher.

To then test your knowledge, you'd run the following command:

```bash
flasher flash bridge-questions.json
```

Replacing `bridge-questions.json` with the name/path to your JSON flashcard deck.

Flasher will then ask you questions from the deck one at a time, waiting for you to press return to show the answer and move on. It's simple, quick and easy.

Flasher doesn't try to check if you got an answer right. Since I (and many other people) don't use flashcards with simple one word answers, this would be crazy hard to implement in a user friendly way. 


## TODO

There's a few things I want add to flasher, namely:

- Add a flashcard creation command. 
- Add the ability to attach an image to a flashcard. This'd likely be implemented by specifying a path/URL to an image in the JSON file and then loading that image with the system's default image loading application. Alternatively, we could try and convert the image to ASCII and then display it in the terminal. That'd be cool but probably kind'a ugly. 
- Add some basic stat tracking functionality. Just track which questions you get right/wrong (using an honour system) and add an option to just ask incorrect ones.

