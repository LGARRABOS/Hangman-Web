package piscine

import (
    "fmt"
)


func PrintHangmanError(nbbeforelose int, lattempts *int, letter string, won *string) {
    if *lattempts != nbbeforelose {
        fmt.Println("Not present in the word, ", nbbeforelose ,"attempts remaining")
        *won = "You have played \"" + letter + "\", that not in the word"
        *lattempts = nbbeforelose
    }
    if nbbeforelose == 0 {
        fmt.Println("The poor José is dead because of you.")
    }
}
