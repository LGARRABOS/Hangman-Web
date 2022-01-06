package piscine

func Affichagefind(word string, letterchoose string, tabunderscore []rune) []rune {
	if  tabunderscore[0] == 0 {
		for i := 0; i < len(tabunderscore); i++ {
			tabunderscore[i] = '_'
		}
	}
	for i:= 0; i < len(tabunderscore); i++ {
		if tabunderscore[i] <= 'z' && tabunderscore[i] >= 'a' {
		} else if tabunderscore[i] == '_' && word[i] == letterchoose[0] {
			tabunderscore[i] = rune(letterchoose[0])
		} else {
			tabunderscore[i] = '_'
		}
	}
	return tabunderscore
}


