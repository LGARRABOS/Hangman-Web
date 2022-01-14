package hangman

import (
	"fmt"
	piscine "hangman/function"
)

func Hangman(letter string, start bool, w string) (string, int, bool, string) {
	attempts := 10
	lattempts := attempts
	c := 0
	var stock []byte
	word := ""
	asciiart := ""
	hidden_word := ""
	won := ""
	fmt.Println(letter)
	if !start {
		attempts, word, stock, asciiart = piscine.Decod()
		fmt.Println("Welcome back, you have", attempts, " attempts.")
	} else {
		fmt.Println("Good luck, you have", attempts, " attempts.")
		word = piscine.RandomWord(w)
		fmt.Println(word)
	}
	tabunderscore := make([]rune, len(word))
	asciiart = ""
	if !start {
		for i := 0; i < len(stock); i++ {
			tabunderscore = piscine.Affichagefind(word, string(stock[i]), tabunderscore)
		}
		hidden_word = piscine.LetterType(tabunderscore, asciiart)
		piscine.PrintHangmanError(attempts, &lattempts, letter, &won)
	} else {
		baseletter := piscine.LetterRandom(word, &stock)
		tabunderscore = piscine.Affichagefind(word, baseletter, tabunderscore)
		hidden_word = piscine.LetterType(tabunderscore, asciiart)
		piscine.PrintHangmanError(attempts, &lattempts, letter, &won)
		piscine.Encod(attempts, word, stock, asciiart)
		return hidden_word, attempts, false, won
	}
	c = 0
	if piscine.AllVerif(letter, &stock, word, &won) {
		if piscine.Verif_letter_in_word(word, letter) {
			tabunderscore = piscine.Affichagefind(word, letter, tabunderscore)
			hidden_word = piscine.LetterType(tabunderscore, asciiart)
			piscine.PrintHangmanError(attempts, &lattempts, letter, &won)
			for i := 0; i < len(word); i++ {
				if rune(word[i]) == tabunderscore[i] {
					c++
				}
			}
		} else {
			attempts--
			hidden_word = piscine.LetterType(tabunderscore, asciiart)
			piscine.PrintHangmanError(attempts, &lattempts, letter, &won)
		}
		if c == len(word) {
			fmt.Println("Congrats !")
			return hidden_word, attempts, true, won
		}
	} else if len(letter) == len(word) {
		if piscine.Complet_word(word, letter) {
			for i := 0; i < len(word); i++ {
				tabunderscore[i] = rune(word[i])
			}
			hidden_word = piscine.LetterType(tabunderscore, asciiart)
			piscine.PrintHangmanError(attempts, &lattempts, letter, &won)
			fmt.Println("\n")
			fmt.Println("Congrats !")
			return hidden_word, attempts, true, won
		} else {
			attempts -= 2
			hidden_word = piscine.LetterType(tabunderscore, asciiart)
			piscine.PrintHangmanError(attempts, &lattempts, letter, &won)
		}
	}
	start = false
	piscine.Encod(attempts, word, stock, asciiart)
	return hidden_word, attempts, false, won
}
