package Jeux

import (
	"fmt"
	"math/rand/v2"
)

var LosePoint int = 0
var LosePointMax int
var Inconnu []string
var Begin bool
var LettreError []string

func Play() {
	LosePoint = 0
	LosePointMax = 11
	var choix string
	var SideWord string
	var Word string = ""
	Inconnu = []string{}
	LettreError = []string{}
	Word = ListeWord[rand.IntN(len(ListeWord))]
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
		fmt.Println("/=========================H.A.N.G.M.A.N===========================\\")
		fmt.Println("                                                               ")
		fmt.Print("   Side this word : ")
		for _, k := range Inconnu {
			fmt.Print(k)
		}
		fmt.Println("")
		fmt.Print("   Nb Error : ", LosePoint)
		fmt.Println("")
		fmt.Print("   Your fault : ")
		for _, L := range LettreError {
			fmt.Print(L, " ")
		}
		fmt.Println("")
		Error()
		fmt.Println("")
		fmt.Println("")
		fmt.Println("      1. Side a word ")
		fmt.Println("      2. Side a letter")
		fmt.Println("      0. Exit")
		fmt.Println(" ")
		fmt.Println("\\=========================N.A.M.G.N.A.H===========================/")
		fmt.Println(" ")
		fmt.Print("Your choice : ")
		fmt.Scanln(&choix)
		switch choix {
		case "1":
			TrouverMot(Word)
			clearScreen()
		case "2":
			TrouverLettre(Word)
			clearScreen()
		default:
			clearScreen()
			break
		}
		if choix == "0" {
			clearScreen()
			Hangman()
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
	clearScreen()
	fmt.Println("/==============H.A.N.G.M.A.N==============\\")
	fmt.Println("")
	fmt.Println("              ", LightRed, "GAME  OVER", Reset)
	fmt.Println("")
	fmt.Println("            Word : ", mot)
	fmt.Println("")
	fmt.Println("     1. New Game    2. Menu    3. Exit")
	fmt.Println("")
	fmt.Println("\\==============N.A.M.G.N.A.H==============/")
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
	clearScreen()
	fmt.Println("/==============H.A.N.G.M.A.N==============\\")
	fmt.Println("")
	fmt.Println("                  ", LightGreen, "WIN", Reset)
	fmt.Println("")
	fmt.Println("            Word : ", mot)
	fmt.Println("")
	fmt.Println("     1. New Game    2. Menu    3. Exit")
	fmt.Println("")
	fmt.Println("\\==============N.A.M.G.N.A.H==============/")
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
