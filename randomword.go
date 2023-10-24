package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

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

	num := randomNumb(array)
	wordtofind := array[num-1]
	wordtofind = "\n" + wordtofind
	return wordtofind
}

func randomNumb(array []string) int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := len(array)
	return (rand.Intn(max-min+1) + min)
}
