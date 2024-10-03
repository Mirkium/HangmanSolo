package Jeux

import "fmt"

func TrouverLettre(mot string) {
	var lettre string
	var choix string
	var Error int = 0
	fmt.Println("================H.A.N.G.M.A.N================")
	fmt.Println("")
	fmt.Println("1. Choose a letter ( Be careful! no exit before entering the letter).")
	fmt.Println("0. Exit")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	if len(choix) == 1 {
		switch choix {
		case "1":
			clearScreen()
			fmt.Println("================H.A.N.G.M.A.N================")
			fmt.Println("")
			fmt.Print("the letter : ")
			fmt.Scanln(&lettre)
			for _, L := range LettreError {
				if lettre == L {
					clearScreen()
					fmt.Println("Not secoond try !")
					fmt.Println("")
					TrouverLettre(mot)
				}
			}
			if Verif(lettre) {
				for i, k := range mot {
					if lettre == string(k) {
						Inconnu[i] = lettre
					} else {
						Error++
					}
				}
				if Error == len(mot) {
					LettreError = append(LettreError, lettre)
					LosePoint++
					if LosePoint > LosePointMax {
						LosePoint = LosePointMax
					}
				}
			} else {
				clearScreen()
				fmt.Println("I want a one letter !")
				fmt.Println("")
				TrouverLettre(mot)
			}
		case "0":
			fmt.Println("Exit")
		default:
			clearScreen()
			TrouverLettre(mot)
		}
	} else {
		TrouverMot(mot)
	}
}

func TrouverMot(mot string) {
	var choix string
	fmt.Println("================H.A.N.G.M.A.N================")
	fmt.Println("")
	fmt.Println("1. Choose a word ( Be careful! no exit before entering the word).")
	fmt.Println("0. Exit")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		var word string
		fmt.Print("Choose a word : ")
		fmt.Scanln(&word)
		if len(word) <= len(mot) {
			for _, L := range LettreError {
				if word == L {
					clearScreen()
					fmt.Println("Not second try")
					fmt.Println("")
					TrouverMot(mot)
				}
			}
			if Verif(word) {
				if word == mot {
					for i, k := range mot {
						Inconnu[i] = string(k)
					}
					Win(mot)
				} else {
					LettreError = append(LettreError, word)
					LosePoint += 2
					if LosePoint > LosePointMax {
						LosePoint = LosePointMax
					}
				}
			} else {
				clearScreen()
				fmt.Println("I want a one word !")
				fmt.Println("")
				TrouverMot(mot)
			}
		} else {
			clearScreen()
			fmt.Println("Retry")
			fmt.Println("")
			LosePoint += 2
			if LosePoint > LosePointMax {
				LosePoint = LosePointMax
				Lose(mot)
			} else {
				TrouverMot(mot)
			}
		}
	case "0":
		fmt.Println("Exit")
		clearScreen()
	default:
		clearScreen()
		TrouverMot(mot)
	}
}
