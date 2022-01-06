package piscine 

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)
func RandomWord() string {
	rand.Seed(time.Now().UnixNano())
	var wordtab []string
	var word string
	f, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wordtab = append(wordtab, scanner.Text())
    }
	word = wordtab[rand.Intn(len(wordtab))]
	return word
	
}

func LetterRandom(mot string, stock *[]byte) string {
	var letter string
	letter = string(mot[rand.Intn(len(mot))])
	*stock = append(*stock, byte(letter[0]))
	return letter 
}
