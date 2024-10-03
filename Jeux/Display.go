package Jeux

import (
	"fmt"
	"io/ioutil"
	"strconv"
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

func Error() {
	File := strconv.Itoa(LosePoint)
	ErrorNb, err := ioutil.ReadFile("Jeux/Display/Display" + File + ".txt") // lire le fichier text.txt
	if LosePoint != 0 {
		if err != nil {
			fmt.Println(err)
		}
	}
	switch LosePoint {
	case 1:
		fmt.Println(LightGreen, string(ErrorNb), Reset)
	case 2:
		fmt.Println(LightGreen, string(ErrorNb), Reset)
	case 3:
		fmt.Println(DarkGreen, string(ErrorNb), Reset)
	case 4:
		fmt.Println(LightYellow, string(ErrorNb), Reset)
	case 5:
		fmt.Println(LightYellow, string(ErrorNb), Reset)
	case 6:
		fmt.Println(OrangeYellow, string(ErrorNb), Reset)
	case 7:
		fmt.Println(LightOrange, string(ErrorNb), Reset)
	case 8:
		fmt.Println(LightOrange, string(ErrorNb), Reset)
	case 9:
		fmt.Println(PinkRed, string(ErrorNb), Reset)
	case 10:
		fmt.Println(LightRed, string(ErrorNb), Reset)
	case 11:
		fmt.Println(LightRed, string(ErrorNb), Reset)
	}
}
