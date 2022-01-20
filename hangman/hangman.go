package hangman

import (
	"fmt"
	piscine "hangman/function"
)

func Hangman(letter string, start bool, w string, user string, diff string) (string, int, bool, string, string, string) {
	attempts := 10
	lattempts := attempts
	c := 0
	var stock []byte
	word := ""
	hidden_word := ""
	won := ""
	used := ""
	fmt.Println(letter)
	if !start {
		attempts, word, stock = piscine.Decod()
		fmt.Println("Welcome back, you have", attempts, " attempts.")
	} else {
		fmt.Println("Good luck, you have", attempts, " attempts.")
		word = piscine.RandomWord(w)
		fmt.Println(word)
	}
	tabunderscore := make([]rune, len(word))
	if !start {
		for i := 0; i < len(stock); i++ {
			tabunderscore = piscine.Affichagefind(word, string(stock[i]), tabunderscore)
		}
		hidden_word = piscine.LetterType(tabunderscore)
		piscine.PrintHangmanError(attempts, &lattempts, letter, &won)
	} else {
		baseletter := piscine.LetterRandom(word, &stock)
		tabunderscore = piscine.Affichagefind(word, baseletter, tabunderscore)
		hidden_word = piscine.LetterType(tabunderscore)
		piscine.PrintHangmanError(attempts, &lattempts, letter, &won)
		piscine.Encod(attempts, word, stock)
		used = piscine.TabToString(stock)
		var userlist []string
		var attlist []int
		var wordlist []string
		userlist = append(userlist, user)
		attlist = append(attlist, attempts)
		wordlist = append(wordlist, word)
		return hidden_word, attempts, false, won, used, word
	}
	c = 0
	if piscine.AllVerif(letter, &stock, word, &won) {
		if piscine.Verif_letter_in_word(word, letter) {
			tabunderscore = piscine.Affichagefind(word, letter, tabunderscore)
			hidden_word = piscine.LetterType(tabunderscore)
			piscine.PrintHangmanError(attempts, &lattempts, letter, &won)
			for i := 0; i < len(word); i++ {
				if rune(word[i]) == tabunderscore[i] {
					c++
				}
			}
		} else {
			attempts--
			hidden_word = piscine.LetterType(tabunderscore)
			piscine.PrintHangmanError(attempts, &lattempts, letter, &won)
		}
		if c == len(word) {
			fmt.Println("Congrats !")
			return hidden_word, attempts, true, won, used, word
		}
	} else if len(letter) == len(word) {
		if piscine.Complet_word(word, letter) {
			for i := 0; i < len(word); i++ {
				tabunderscore[i] = rune(word[i])
			}
			hidden_word = piscine.LetterType(tabunderscore)
			piscine.PrintHangmanError(attempts, &lattempts, letter, &won)
			fmt.Println("\n")
			fmt.Println("Congrats !")
			return hidden_word, attempts, true, won, used, word
		} else {
			attempts -= 2
			hidden_word = piscine.LetterType(tabunderscore)
			piscine.PrintHangmanError(attempts, &lattempts, letter, &won)
		}
	}
	start = false
	piscine.Encod(attempts, word, stock)
	used = piscine.TabToString(stock)
	return hidden_word, attempts, false, won, used, word
}
