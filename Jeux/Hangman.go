package Jeux

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var ListeWord []string

func Hangman() {
	var choix string
	fmt.Println("       H-A-N-G-M-A-N ")
	fmt.Println("======[]==============|========")
	fmt.Println("      []            __|__")
	fmt.Println("      []            |   |")
	fmt.Println("      []            |___|")
	fmt.Println("      []              |")
	fmt.Println("      []             /|\\")
	fmt.Println("      []            / | \\")
	fmt.Println("      []             / \\")
	fmt.Println("      []            /   \\")
	fmt.Println("     /[]\\")
	fmt.Println(" ___/_[]_\\____")
	fmt.Println(" ")
	fmt.Println("Jeux/Display/Display" + string(LosePoint) + ".txt")
	fmt.Println(" 1. Play    2. Langage   3. Rules    4. Exit")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		clearScreen()
		Play()
	case "2":
		clearScreen()
		Langage()
		clearScreen()
		Hangman()
	case "3":
		clearScreen()
		Rules()
		clearScreen()
		Hangman()
	case "4":
		break
	default:
		clearScreen()
		Hangman()
	}
}

func clearScreen() {
	var cmd *exec.Cmd

	// Détecter le système d'exploitation
	switch runtime.GOOS {
	case "windows":
		// Commande pour Windows
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		// Commande pour les systèmes Unix-like
		cmd = exec.Command("clear")
	}

	// Définir la sortie de la commande sur Stdout
	cmd.Stdout = os.Stdout
	cmd.Run()
}
