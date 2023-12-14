package hangman

import (
	"flag"
	"fmt"
	"os"
)

type HangManData struct {
	Word             string  // Word composed of '_', ex: H_ll_
	ToFind           string  // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int     // Number of attempts left
	HangmanPositions [10]int // It can be the array where the positions parsed in "hangman.txt" are stored
	Letters          string
	LetterInput      string
	File             string
	TypeOfGame       bool
}

// Create the struct to start the game
func New(file string, ASCII_File string) *HangManData {
	word_file := "dic/" + file
	var H HangManData
	H.ToFind = RandomWord(string((word_file)))
	H.Word = RandomWordUnderscore(H.ToFind)
	H.Attempts = 10
	H.HangmanPositions = [10]int{72, 64, 56, 48, 40, 32, 24, 16, 8, 0}
	H.File = ASCII_File
	H.TypeOfGame = true
	H.LetterInput = " "
	return &H
}

func Test() {
	fmt.Print("Hello")
}

// Detect with flag if the user want start a new game or load a game
func HangmanSolver() {
	flag.String("startWith", "default", "File name to start with")
	flag.String("letterFile", "default", "File name to choose ASCII")
	flag.Parse()
	if len(os.Args[1:]) >= 2 {
		if os.Args[1] == "--startWith" && os.Args[2] == "save.txt" {
			var H HangManData
			LoadGame("save.txt", &H)
			if !H.TypeOfGame {
				StandardHangmanGame(&H)
			} else {
				Box(&H)
			}
		} else if os.Args[1] == "--letterFile" && (os.Args[2] == "standard.txt" || os.Args[2] == "shadow.txt" || os.Args[2] == "thinkertoy.txt" || os.Args[2] == "default.txt") {
			state := Menu()
			H := New(os.Args[3], os.Args[2])
			FirstLetter(H)
			if !state {
				H.TypeOfGame = false
				StandardHangmanGame(H)
			} else {
				H.TypeOfGame = true
				Box(H)
			}
		}
	} else if len(os.Args[1:]) == 1 {
		state := Menu()
		H := New(os.Args[1], "default.txt")
		FirstLetter(H)
		if !state {
			H.TypeOfGame = false
			StandardHangmanGame(H)
		} else {
			H.TypeOfGame = true
			Box(H)
		}
	} else {
		fmt.Println("Syntax problem")
		os.Exit(6)
	}
}

func FirstLetter(H *HangManData) {
	letteruse := ""
	for _, i := range LettersUse(H) {
		letteruse += i + " | "
	}
	H.Letters = letteruse
}