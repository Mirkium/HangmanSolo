package Jeux

func Verif(s string) bool {
	for _, k := range s {
		if (k >= 'a' && k <= 'z') || (k >= 'A' && k <= 'Z') {
			return true
		}
	}
	return false
}
