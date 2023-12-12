package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Pick a random word in the file
func RandomWord(file string) string {
	array := []string{}
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		array = append(array, fileScanner.Text())
	}

	num := RandomNumbArray(array)
	wordtofind := array[num-1]
	wordtofind = "\n" + wordtofind
	return wordtofind
}

// Transform the word to find with underscore
func RandomWordUnderscore(word string) string {
	wordUnderscore := []string{}
	n := len(word)/2 - 1
	reveal := word[RandomNumbWord(word)-1]
	for i := 1; i <= len(word)-1; i++ {
		if reveal == word[i] {
			wordUnderscore = append(wordUnderscore, string(word[i]))
			n--
		} else {
			wordUnderscore = append(wordUnderscore, "_")
		}
	}
	help := 0
	for {
		if help > 100 {
			n--
		}
		if n >= 1 {
			count := 0
			revealbis := word[RandomNumbWord(word)-1]
			for k := range word {
				if word[k] == revealbis {
					count++
				}
			}
			if count > 1 {
				help++
			} else {
				for j := 0; j <= len(word)-2; j++ {
					if revealbis == word[j+1] && wordUnderscore[j] == "_" {
						wordUnderscore[j] = string(word[j+1])
						n--
					}
				}
			}
		} else {
			finalWord := ""
			for _, letters := range wordUnderscore {
				finalWord += " " + letters
			}
			return finalWord
		}
	}
}

// Pick a random number
func RandomNumbArray(array []string) int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := len(array)
	return (rand.Intn(max-min+1) + min)
}

func RandomNumbWord(word string) int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := len(word)
	return (rand.Intn(max-min+1) + min)
}
