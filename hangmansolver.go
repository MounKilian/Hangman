package hangman

import (
	"flag"
	"fmt"
	"os"
)

type HangManData struct {
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]int    // It can be the array where the positions parsed in "hangman.txt" are stored
	Letters          string     // List of all letters used
	LetterInput      string     // LetterInput by the user
	File             string     // File for ASCII Art
	TypeOfGame       bool       // True if the game is with tview and false if is it without tview
	WordFile         string     // To know wich word file choose the user
	Username         string     // To know the username of the user (HangmanWeb)
	Scoreboard       [][]string // To store the scoreboard of the game (HangmanWeb)
	NewScore         []string   // To store the score of the user (HangmanWeb)
	Point            int        // Points of the user during the game (HangmanWeb)
	Level            string     // To know if user is in hard, medium or easy level (HangmanWeb)
}

// Create the struct to start the game
func New(file string, ASCII_File string) *HangManData {
	var H HangManData
	H.WordFile = "dic/" + file
	H.ToFind = RandomWord(string((H.WordFile)))
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
