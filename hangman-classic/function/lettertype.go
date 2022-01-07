package piscine

import (
	"fmt"
)
func LetterType(tabunderscore []rune, asciiart string) {
	if asciiart != "" {
		Ascii_art_min(tabunderscore, asciiart)
	} else {
		for i := 0; i < len(tabunderscore); i++ {
			fmt.Printf(string(tabunderscore[i]))
			fmt.Printf(" ")
		}
		fmt.Println()
	}
}

func TransfoLetter(tabunderscore []rune) []int {
	nletter := make([]int, len(tabunderscore))
	for i := 0; i < len(tabunderscore); i++ {
		for j := 97; j < 123; j++ {
			if tabunderscore[i] == rune(j) {
				nletter[i] = 587+ 9*(j-97)
			}
		}
		if tabunderscore[i] == '_' {
			nletter[i] = 569
		}
	}
	return nletter
}