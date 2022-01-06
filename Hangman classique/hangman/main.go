package main

import (
	"fmt"
	"os"
	"piscine"
)

func main() {
	attempts := 10
	lattempts := attempts
	var choice string
	c := 0
	var stock []byte
	word := ""
	asciiart := ""
	if len(os.Args) == 3 && os.Args[1] == "--startWith" {
		attempts, word, stock, asciiart = piscine.Decod()
		lattempts = attempts
		fmt.Println("Welcome back, you have", attempts, " attempts.")
	} else {
		fmt.Println("Good luck, you have", attempts, " attempts.")
		word = piscine.RandomWord()
	}
	tabunderscore := make([]rune, len(word))
	if len(os.Args) == 4 && os.Args[2] == "--letterFile" {
		asciiart = os.Args[3]
	}
	if len(os.Args) == 3 && os.Args[1] == "--startWith" {
		for i := 0; i < len(stock); i++ {
			tabunderscore = piscine.Affichagefind(word, string(stock[i]), tabunderscore)
		}
		piscine.LetterType(tabunderscore, asciiart)
		piscine.PrintHangmanError(attempts, &lattempts)
	} else {
		baseletter := piscine.LetterRandom(word, &stock)
		tabunderscore = piscine.Affichagefind(word, baseletter, tabunderscore)
		piscine.LetterType(tabunderscore, asciiart)
		piscine.PrintHangmanError(attempts, &lattempts)
	}
	for attempts > 0 {
		choice = piscine.Choice()
		if choice == "STOP" {
			piscine.Encod(attempts, word, stock, asciiart)
			break

		}
		c = 0
		if piscine.AllVerif(choice, &stock, word) {
			if piscine.Verif_letter_in_word(word, choice) {

				tabunderscore = piscine.Affichagefind(word, choice, tabunderscore)
				piscine.LetterType(tabunderscore, asciiart)
				piscine.PrintHangmanError(attempts, &lattempts)
				for i := 0; i < len(word); i++ {
					if rune(word[i]) == tabunderscore[i] {
						c++
					}
				}
			} else {
				attempts--
				piscine.LetterType(tabunderscore, asciiart)
				piscine.PrintHangmanError(attempts, &lattempts)
			}
			if c == len(word) {
				fmt.Println("Congrats !")
				break
			}
		} else if len(choice) == len(word) {
			if piscine.Complet_word(word, choice) {
				for i := 0; i < len(word); i++ {
					tabunderscore[i] = rune(word[i])
				}
				piscine.LetterType(tabunderscore, asciiart)
				piscine.PrintHangmanError(attempts, &lattempts)
				fmt.Println("\n")
				fmt.Println("Congrats !")
				break
			} else {
				attempts -= 2
				piscine.LetterType(tabunderscore, asciiart)
				piscine.PrintHangmanError(attempts, &lattempts)
			}
		}
	}
}
