package hangman

import (
	"fmt"
	"time"
)

func StandardHangmanGame(file string, H *HangManData) {
	NewTextStandard(file, H)
	for {
		fmt.Println("-------------------------- LETTER CHOICE -------------------------")
		fmt.Print("Choose : ")
		fmt.Scanln(&H.LetterInput)
		fmt.Println("------------------------------------------------------------------")
		if H.LetterInput == "STOP" {
			Save(H)
			break
		}
		//If the user enter a letter
		if !VerifIfAlreadyUse(H) && (H.LetterInput >= "a" && H.LetterInput <= "z") {
			if len(H.LetterInput) == 1 {
				Verification(H)
				NewTextStandard(file, H)
				if WordFind(H) {
					time.Sleep(1 * time.Second)
					Victory(H)
					break
				}
				//If the user enter a word
			} else if len(H.LetterInput) > 1 {
				win := EnterWord(H)
				NewTextStandard(file, H)
				if win {
					Victory(H)
					break
				}
			}
			//If the user enter an invalid or already use letter or word
		} else {
			NewTextStandard(file, H)
		}
		if H.Attempts <= 0 {
			time.Sleep(1 * time.Second)
			Defeat(H)
			break
		}
	}
}

func NewTextStandard(file string, H *HangManData) {
	H.Letters += H.LetterInput + " | "
	fmt.Println("----------------------- ATTEMPTS REMAINING -----------------------")
	fmt.Println(H.Attempts)
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("------------------------- HANGMAN STATE --------------------------")
	fmt.Println(HangmanState(H))
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("-------------------------- LETTERS USE ---------------------------")
	fmt.Println(H.Letters)
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("------------------------------ WORD ------------------------------")
	help := H.Word
	ConvertToASCII("ASCII/"+file, H)
	fmt.Println(H.Word)
	H.Word = help
	fmt.Println("------------------------------------------------------------------")
}
