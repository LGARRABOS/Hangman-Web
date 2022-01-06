package piscine

import (
	"bufio"
	"os"
	"fmt"
)
func Choice() string {
	scanner := bufio.NewScanner(os.Stdin)
	var choice string
	fmt.Printf("Choose: ")
	for scanner.Scan() {
		choice = choice + scanner.Text()
		break
  	}
	return choice
	
}