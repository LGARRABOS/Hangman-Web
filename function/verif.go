package piscine

func Verif_letter_in_word(word string, choice string) bool {
	for i := 0; i < len(word); i++ {
		if word[i] == choice[0] {
			return true
		}
	}
	return false
}

func Verif_letter(choice string) bool {
	if choice[0] < 123 && choice[0] > 96 {
		return true
	}
	return false
}

func Verif_taille(choice string, word string) bool {
	if len(choice) == 1 {
		return true
	}
	return false
}

func Lettre_utiliser(use *[]byte, choice string) bool {
	a := *use
	for i := 0; i < len(*use); i++ {
		if a[i] == choice[0] {
			return false
		}
	}
	*use = append(*use, choice[0])
	return true
}

func Complet_word(word string, choice string) bool {
	for i := 0; i < len(word); i++ {
		if word[i] != choice[i] {
			return false
		}
	}
	return true
}
