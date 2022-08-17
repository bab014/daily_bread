package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Book struct {
	Book     string     `json:"book"`
	Chapters []Chapters `json:"chapters"`
}
type Verses struct {
	Verse string `json:"verse"`
	Text  string `json:"text"`
}
type Chapters struct {
	Chapter string   `json:"chapter"`
	Verses  []Verses `json:"verses"`
}

// NewBook function to create a new book
// TODO: figure out how to get this from json
func NewBook(f *os.File) *Book {
	// read from the already opened file
	// and check the error
	book, err := io.ReadAll(f)
	if err != nil {
		// TODO: better error handling
		panic(err)
	}
	// create a json decoder
	decoder := json.NewDecoder(bytes.NewReader(book))
	// decode the json into a book struct
	var b Book
	if err = decoder.Decode(&b); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	return &b
}
