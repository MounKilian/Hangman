package hangman

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

func VerifIfAlreadyUse(H *HangManData) bool {
	for _, i := range H.Letters {
		if H.LetterInput == string(i) {
			return true
		}
	}
	return false
}
