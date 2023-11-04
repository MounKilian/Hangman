package hangman

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Encode the current struct of the game in a byte in save.txt
func Save(H *HangManData) {
	save, err := json.Marshal(H)
	if err != nil {
		os.Exit(1)
	}

	file, err := os.Create("save.txt")
	if err != nil {
		os.Exit(2)
	}
	defer file.Close()

	_, err = file.Write(save)
	if err != nil {
		os.Exit(3)
	}
}

// Decode the struct in the save.txt
func LoadGame(file string, H *HangManData) {
	load, err := ioutil.ReadFile(file)
	if err != nil {
		os.Exit(4)
	}
	err = json.Unmarshal(load, H)
	if err != nil {
		os.Exit(5)
	}
}
