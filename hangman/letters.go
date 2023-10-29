package hangman

// Give the letters used in the first reveal
func LettersUse(H *HangManData) []string {
	count := 0
	lettersuse := []string{}
	for _, i := range H.Word {
		if i >= 'a' && i <= 'z' {
			for _, k := range lettersuse {
				if k == string(i) {
					count++
				}
			}
			if count < 1 {
				lettersuse = append(lettersuse, string(i))
			}
		}
		count = 0
	}
	return lettersuse
}

// Give the new information when a letter is enter in the user input
func Verification(H *HangManData) {
	new_word := ""
	array := []string{}
	count := 0
	for _, k := range H.Word {
		if k != ' ' {
			array = append(array, string(k))
		}
	}
	for i := 1; i <= len(H.ToFind)-1; i++ {
		if string(H.ToFind[i]) == H.LetterInput {
			array[i-1] = H.LetterInput
			count++
		}
	}
	for _, letter := range array {
		new_word += " " + letter
	}
	if count == 0 {
		H.Attempts--
	}
	H.Word = ""
	H.Word = new_word
}

// Verify if the letter or word enter in the user input is already used
func VerifIfAlreadyUse(H *HangManData) bool {
	word := ""
	for _, i := range H.Letters {
		if string(i) != " " && string(i) != "|" {
			word += string(i)
		} else if string(i) == "|" {
			if H.LetterInput == word {
				return true
			}
			word = ""
		}
	}
	return false
}

// Give the new information when a word is enter in the user input and know if the word is complete
func EnterWord(H *HangManData) bool {
	new_word := ""
	count := 0
	etat := true
	if len(H.LetterInput) == len(H.ToFind)-1 {
		for i := 0; i <= len(H.LetterInput)-1; i++ {
			if string(H.LetterInput[i]) != string(H.ToFind[i+1]) {
				count++
			}
		}
		if count >= 1 {
			etat = false
		}
	} else {
		etat = false
	}
	if !etat {
		H.Attempts -= len(H.LetterInput)
	} else {
		for _, k := range H.ToFind {
			new_word += " " + string(k)
		}
		H.Word = new_word
	}
	return etat
}

// Return true if the word is complete false if isn't
func WordFind(H *HangManData) bool {
	count := 0
	etat := true
	new_word := ""
	for _, j := range H.Word {
		if j != ' ' {
			new_word += string(j)
		}
	}
	if len(new_word) == len(H.ToFind)-1 {
		for i := 0; i <= len(new_word)-1; i++ {
			if string(new_word[i]) != string(H.ToFind[i+1]) {
				count++
			}
		}
		if count >= 1 {
			etat = false
		}
	} else {
		etat = false
	}
	return etat
}
