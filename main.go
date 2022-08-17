package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
)

const REPO_PATH = "./Bible-kjv-master"

func main() {
	r1 := newRandSeed()

	b, err := os.Open(path.Join(REPO_PATH, randomBook(r1)))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer b.Close()

	content, err := io.ReadAll(b)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var j *Book
	err = json.Unmarshal(content, &j)

	rc := randomChapter(j, r1)
	rv := randomVerse(rc, r1)
	fmt.Println(j.Book, rc.Chapter, rv.Verse, rv.Text)
}
