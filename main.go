package main

import (
	game "Hangman/Jeux"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("C'est quand la pause branlette ?")
	} else {
		Data, err := ioutil.ReadFile("Global.txt") // lire le fichier text.txt
		if err != nil {
			fmt.Println(err)
		}
		game.ListeWord = []string{}
		ligneTexte := string(Data)
		var word string
		for _, k := range ligneTexte {
			if k != '\n' && k != '\r' { // Utiliser && au lieu de ||
				word += string(k)
			} else if word != "" { // Ajouter le mot si non vide
				game.ListeWord = append(game.ListeWord, word)
				word = ""
			}
		}
		game.Hangman()
	}
	
}
