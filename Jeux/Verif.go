package Jeux


func Verif(s string) bool {
	for _, k := range s {
		if (k >= 'a' && k <= 'z') || (k >= 'A' && k <= 'Z') || (k >= 'ç' && k <= 'ï') || (k >= 'ñ' && k <= 'ö') || (k >= 'ù' && k <= 'ý') || (k >= 'ß' && k <= 'å') || (k >= 'Ù' && k <= 'Ý') || (k >= 'Ñ' && k <= 'Ö') || (k >= 'Ç' && k <= 'Ï') || (k >= 'À' && k <= 'Å') {
			return true
		} 
	}
	return false
}
