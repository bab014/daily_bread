package main

import (
	"os"
	"testing"
)

// testing if randomBook doesn't return an empty string
func TestRandomBook(t *testing.T) {
	// get the list of possible books
	pbooks, err := os.ReadDir(REPO_PATH)
	if err != nil {
		t.Error(err)
	}

	// get a random book formatted to match the json file name
	book := randomBook()
	// book := "notabook.json"

	// loop through pbooks and check if the book is in the list
	for _, pbook := range pbooks {
		if pbook.Name() == book {
			return
		}
	}

	// if we get here, the book is not in the list
	t.Errorf("randomBook() returned %s. Not a valid book\n", book)
}
