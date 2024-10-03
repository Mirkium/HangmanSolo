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
	fmt.Println("/================H.A.N.G.M.A.N================\\")
	fmt.Println("")
	fmt.Println("       ======[]==============|========")
	fmt.Println("             []            __|__")
	fmt.Println("             []            |   |")
	fmt.Println("             []            |___|")
	fmt.Println("             []              |")
	fmt.Println("             []             /|\\")
	fmt.Println("             []            / | \\")
	fmt.Println("             []             / \\")
	fmt.Println("             []            /   \\")
	fmt.Println("            /[]\\")
	fmt.Println("        ___/_[]_\\____")
	fmt.Println(" ")
	fmt.Println("        1. Play    2. Rules    3. Exit")
	fmt.Println("")
	fmt.Println("\\================N.A.M.G.N.A.H================/")
	fmt.Print("Your choice : ")
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		clearScreen()
		Play()
	case "2":
		clearScreen()
		Rules()
		clearScreen()
		Hangman()
	case "3":
		clearScreen()
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
