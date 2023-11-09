package hangman

import (
	"fmt"
	"strconv"
	"time"
)

func StandardHangmanGame(H *HangManData) {
	fmt.Println()
	fmt.Println("----------------------- ATTEMPTS REMAINING -----------------------")
	fmt.Println("Good Luck, " + strconv.Itoa(H.Attempts) + " attempts remaining")
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("------------------------- HANGMAN STATE --------------------------")
	fmt.Println(HangmanState(H))
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("-------------------------- LETTERS USE ---------------------------")
	fmt.Println(H.Letters)
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("------------------------------ WORD ------------------------------")
	if H.File != "default.txt" {
		fmt.Println(ConvertToASCII("ASCII/"+H.File, H))
	} else {
		fmt.Println(H.Word)
	}
	fmt.Println("------------------------------------------------------------------")
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
				NewTextStandard(H)
				if WordFind(H) {
					time.Sleep(1 * time.Second)
					Victory(H)
					break
				}
				//If the user enter a word
			} else if len(H.LetterInput) > 1 {
				win := EnterWord(H)
				NewTextStandard(H)
				if win {
					Victory(H)
					break
				}
			}
			//If the user enter an invalid or already use letter or word
		} else {
			fmt.Println("Letter invalid or already use  ")
		}
		if H.Attempts <= 0 {
			time.Sleep(1 * time.Second)
			Defeat(H)
			break
		}
	}
}

// Refresh the data after all the input of the user
func NewTextStandard(H *HangManData) {
	H.Letters += H.LetterInput + " | "
	fmt.Println()
	fmt.Println("----------------------- ATTEMPTS REMAINING -----------------------")
	fmt.Println("Good Luck, " + strconv.Itoa(H.Attempts) + " attempts remaining")
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("------------------------- HANGMAN STATE --------------------------")
	fmt.Println(HangmanState(H))
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("-------------------------- LETTERS USE ---------------------------")
	fmt.Println(H.Letters)
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("------------------------------ WORD ------------------------------")
	if H.File != "default.txt" {
		fmt.Println(ConvertToASCII("ASCII/"+H.File, H))
	} else {
		fmt.Println(H.Word)
	}
	fmt.Println("------------------------------------------------------------------")
}
