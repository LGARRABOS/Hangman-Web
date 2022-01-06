package piscine

import (
    "fmt"
    "os"
	"bufio"
	"log"
)


func PrintHangmanError(nbbeforelose int, lattempts *int) {
    var wordtab []string
    if *lattempts != nbbeforelose {
      fmt.Println("Not present in the word, ", nbbeforelose ,"attempts remaining")
      *lattempts = nbbeforelose
    }
    f, err := os.Open("hangman.txt")
    if err != nil {
        log.Fatal(err)
    }
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wordtab = append(wordtab, scanner.Text())
    }
    if nbbeforelose == 9 {
        for i := 0; i < 7; i++ {
            fmt.Println(wordtab[i])
        }
    } else if nbbeforelose == 8 {
        for i := 7; i < 15; i++ {
            fmt.Println(wordtab[i])
        }
    } else if nbbeforelose == 7 {
        for i := 15; i < 23; i++ {
            fmt.Println(wordtab[i])
        }
    } else if nbbeforelose == 6 {
        for i := 23; i < 31; i++ {
            fmt.Println(wordtab[i])
        }
    } else if nbbeforelose == 5 {
        for i := 31; i < 39; i++ {
            fmt.Println(wordtab[i])
        }
    } else if nbbeforelose == 4 {
        for i := 39; i < 47; i++ {
            fmt.Println(wordtab[i])
        }
    } else if nbbeforelose == 3 {
        for i := 47; i < 55; i++ {
            fmt.Println(wordtab[i])
        }
    } else if nbbeforelose == 2 {
        for i := 55; i < 63; i++ {
            fmt.Println(wordtab[i])
        }
    } else if nbbeforelose == 1 {
        for i := 63; i < 71; i++ {
            fmt.Println(wordtab[i])
        }
    } else if nbbeforelose == 0 {
        for i := 71; i < 78; i++ {
            fmt.Println(wordtab[i])
        }
        fmt.Println("The poor José is dead because of you.")
    }
}
