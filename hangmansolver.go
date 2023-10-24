package hangman

import "os"

func HangmanSolver() {
	arg := "dic/" + os.Args[1]
	Box(RandomWord(string(arg)))
}
