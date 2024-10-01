package Jeux

import (
	"fmt"
	"io/ioutil"
)

func Langage() {
	WordEnglish, err := ioutil.ReadFile("Jeux/WordEnglish.txt") // lire le fichier text.txt
	if err != nil {
		fmt.Println(err)
	}
	WordFrench, err := ioutil.ReadFile("Jeux/WordFrench.txt")
	if err != nil {
		fmt.Println(err)
	}
	Data, err := ioutil.ReadFile("Jeux/../Global.txt") // lire le fichier text.txt
	if err != nil {
		fmt.Println(err)
	}
	var choix string
	fmt.Println("===============LANGAGE===============")
	fmt.Println("")
	fmt.Println("             1. French")
	fmt.Println("")
	fmt.Println("             2. English")
	fmt.Println("")
	fmt.Println("             3. Global")
	fmt.Println("")
	fmt.Println("             0. Exit")
	fmt.Println("")
	fmt.Println("=====================================")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		ListeWord = []string{}
		ligneTexte := string(WordFrench)
		var word string
		for _, k := range ligneTexte {
			if k != '\n' && k != '\r' { // Utiliser && au lieu de ||
				word += string(k)
			} else if word != "" { // Ajouter le mot si non vide
				ListeWord = append(ListeWord, word)
				word = ""
			}
		}
		// Si le dernier mot n'est pas ajouté (pas de \n à la fin)
		if word != "" {
			ListeWord = append(ListeWord, word)
		}
	case "2":
		ListeWord = []string{}
		ligneTexte := string(WordEnglish)
		var word string
		for _, k := range ligneTexte {
			if k != '\n' && k != '\r' { // Utiliser && au lieu de ||
				word += string(k)
			} else if word != "" { // Ajouter le mot si non vide
				ListeWord = append(ListeWord, word)
				word = ""
			}
		}
		// Si le dernier mot n'est pas ajouté (pas de \n à la fin)
		if word != "" {
			ListeWord = append(ListeWord, word)
		}
	case "3":
		ListeWord = []string{}
		ligneTexte := string(Data)
		var word string
		for _, k := range ligneTexte {
			if k != '\n' && k != '\r' { // Utiliser && au lieu de ||
				word += string(k)
			} else if word != "" { // Ajouter le mot si non vide
				ListeWord = append(ListeWord, word)
				word = ""
			}
		}
	case "0":
		break
	default:
		clearScreen()
		Langage()
	}
}
