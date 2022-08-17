package main

import (
	"encoding/json"
	"io"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
)

// randomBook returns a random book in its json file format
func randomBook(r1 *rand.Rand) string {
	// read in the Books.json file
	b, err := os.Open(path.Join(REPO_PATH, "Books.json"))
	if err != nil {
		panic(err)
	}
	defer b.Close()

	// get all the contents of the file into a variable
	content, err := io.ReadAll(b)
	if err != nil {
		panic(err)
	}

	// unmarshal the books.json file into a []string
	var j []string
	err = json.Unmarshal(content, &j)
	if err != nil {
		panic(err)
	}

	// random number based on the length of the bookds array
	rn := r1.Intn(len(j))

	// format the book name to remove spaces and add the json extension
	book := strings.Replace(j[rn], " ", "", -1) + ".json"

	return book
}

// randomChapter takes in a Book and returns a random chapter
func randomChapter(book *Book, r1 *rand.Rand) *Chapters {
	// random number based on the length of the chapters array
	rn := r1.Intn(len(book.Chapters))

	return &book.Chapters[rn]
}

// randomeVerse takes in a Chapter and returns a random Verses
func randomVerse(chapter *Chapters, r1 *rand.Rand) *Verses {
	// random number based on the length of the verses array
	rn := r1.Intn(len(chapter.Verses))

	return &chapter.Verses[rn]
}

func newRandSeed() *rand.Rand {
	s1 := rand.NewSource(time.Now().UnixNano())
	return rand.New(s1)
}
