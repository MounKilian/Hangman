package hangman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Save(H *HangManData) {
	save, err := json.Marshal(H)
	if err != nil {
		fmt.Println("Erreur lors de la conversion en JSON:", err)
		return
	}

	file, err := os.Create("save.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(save)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
		return
	}
	fmt.Println(string(save))
}

func LoadGame(file string, H *HangManData) {
	jsonData, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier JSON:", err)
		return
	}
	err = json.Unmarshal(jsonData, H)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation JSON:", err)
		return
	}
}
