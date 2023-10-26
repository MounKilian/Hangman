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
