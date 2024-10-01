package Jeux

import (
	"fmt"
	"math/rand/v2"
)

const (
	LightGreen   = "\033[38;5;10m"  // Vert fluo
	DarkGreen    = "\033[38;5;28m"  // Vert foncé
	LightYellow  = "\033[38;5;226m" // Jaune fluo
	OrangeYellow = "\033[38;5;214m" // Jaune orangé
	LightOrange  = "\033[38;5;208m" // Orange fluo
	PinkRed      = "\033[38;5;197m" // Pink red
	LightRed     = "\033[38;5;196m" // Rouge fluo
	Reset        = "\033[0m"        // Réinitialisation de la couleur
)

var LosePoint int
var LosePointMax int
var Inconnu []string
var Begin bool
var LettreError []string

func Play() {
	LosePoint = 0
	LosePointMax = 11
	var choix string
	var SideWord string
	var Word string = ListeWord[rand.IntN(len(ListeWord))]
	Word = ToLower(Word)
	for i := 0; i <= len(Word)-1; i++ {
		Inconnu = append(Inconnu, "_")
	}

	Begin = true
	for {
		for _, k := range Inconnu {
			SideWord += string(k)
		}
		if Begin {
			RandomHelp(Word)
		}
		clearScreen()
		fmt.Println("================H.A.N.G.M.A.N================")
		fmt.Println("")
		fmt.Println("Side this word : ", Inconnu)
		fmt.Println("")
		fmt.Println("Nb Error : ", LosePoint)
		fmt.Println("Your fault : ", LettreError)
		Error()
		fmt.Println("")
		fmt.Println("")
		fmt.Println("1. Side a word ")
		fmt.Println("2. Side a letter")
		fmt.Println("0. Exit")
		fmt.Println(" ")
		fmt.Print("Your choice : ")
		fmt.Scanln(&choix)
		switch choix {
		case "1":
			clearScreen()
			TrouverMot(Word)
			clearScreen()
		case "2":
			clearScreen()
			TrouverLettre(Word)
			clearScreen()
		case "0":
			clearScreen()
			Hangman()
		default:
			clearScreen()
			break
		}
		if SideWord == Word {
			Win(Word)
			break
		} else if LosePoint >= LosePointMax {
			Lose(Word)
			break
		} else {
			continue
		}
	}
}

func RandomHelp(mot string) {
	help := mot[rand.IntN(len(mot))]
	for i, k := range mot {
		if string(help) == string(k) {
			Inconnu[i] = string(help)
		}
	}
	Begin = false
}


func Lose(mot string) {
	var choix string
	fmt.Println("---------------H.A.N.G.M.A.N---------------")
	fmt.Println("")
	fmt.Println("                GAME  OVER")
	fmt.Println("")
	fmt.Println(" Word : ", mot)
	fmt.Println("")
	fmt.Println("   1. New Game    2. Menu    3. Exit")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		clearScreen()
		Play()
	case "2":
		clearScreen()
		Hangman()
	case "3":
		break
	default:
		clearScreen()
		Lose(mot)
	}
}

func Win(mot string) {
	var choix string
	fmt.Println("---------------H.A.N.G.M.A.N---------------")
	fmt.Println("")
	fmt.Println("                   WIN")
	fmt.Println("")
	fmt.Println(" Word : ", mot)
	fmt.Println("")
	fmt.Println("   1. New Game    2. Menu    3. Exit")
	fmt.Println("")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		clearScreen()
		Play()
	case "2":
		clearScreen()
		Hangman()
	case "3":
		break
	default:
		clearScreen()
		Lose(mot)
	}
}

func ToLower(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			result += string(char + 32)
		} else {
			result += string(char)
		}
	}
	return result
}
