package piscine

import (
	"fmt"
)

func AllVerif(choice string, stock *[]byte, word string, won *string) bool {	
	if Verif_taille(choice, word) {
		if Verif_letter(choice) {
			if Lettre_utiliser(&*stock, choice) {
				return true
			} else {
				fmt.Println("The letter given to have already been proposed")
				*won = "The letter given to have already been proposed"
			}
		} else {
			fmt.Println("Your argument can only contain a lowercase letter")
			*won = "Your argument can only contain a lowercase letter"
		}
	} else {
		fmt.Println(choice, " ", word)
		if len(word) != len(choice) {
			fmt.Println("Your argument contains too many letters or not enough")
			*won = "Your argument contains too many letters or not enough"
		}
	}
	return false
}
