package piscine

import (
    "fmt"
)


func PrintHangmanError(nbbeforelose int, lattempts *int) {
    if *lattempts != nbbeforelose {
      fmt.Println("Not present in the word, ", nbbeforelose ,"attempts remaining")
      *lattempts = nbbeforelose
    }
    if nbbeforelose == 0 {
        fmt.Println("The poor José is dead because of you.")
    }
}
