package piscine

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

func Ascii_art_min(tabunderscore []rune, asciiart string) {
	var wordtab []string
	nletter := make([]int, len(tabunderscore))
	f, err := os.Open(asciiart)
    if err != nil {
        log.Fatal(err)
    }
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wordtab = append(wordtab, scanner.Text())
    }
	nletter = TransfoLetter(tabunderscore)

	for i := -1; i <= 7; i++ {
		for j := 0; j < len(nletter); j++ {
			fmt.Printf(wordtab[nletter[j]+i])
		}
		fmt.Println()
	}
	
}