package main

import (
	"hangman"
)

func main() {
	hangman.Box(hangman.SearchWord("words.txt"))
}
