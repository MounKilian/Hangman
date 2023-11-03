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
}

func New() *HangManData {
	arg := "dic/" + os.Args[1]
	var H HangManData
	H.ToFind = RandomWord(string(arg))
	H.Word = RandomWordUnderscore(H.ToFind)
	H.Attempts = 10
	H.HangmanPositions = [10]int{72, 64, 56, 48, 40, 32, 24, 16, 8, 0}
	return &H
}

func HangmanSolver() {
	flag.String("startWith", "default", "Specify a file name to start with")
	flag.Parse()
	if len(os.Args[1:]) == 2 {
		if os.Args[2] == "save.txt" {
			var H HangManData
			LoadGame("save.txt", &H)
			Box(&H)
		}
	} else if len(os.Args[1:]) == 1 {
		H := New()
		letteruse := ""
		for _, i := range LettersUse(H) {
			letteruse += i + " | "
		}
		H.Letters = letteruse
		Box(H)
	} else {
		fmt.Print("Syntax problem")
	}
}
