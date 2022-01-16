package piscine

import (
	"strconv"
)
func TabToString(stock []byte) string {
	strstock := ""
	for i := 0; i < len(stock); i++ {
		strstock = strstock + string(stock[i]) + ", "
	}
	return strstock
}

func AddScore(user string, attempts int, word string) ([]string, []int, []string) {
	userlist, attlist, wordlist := DecodSB()
	userlist = append(userlist, user)
	attlist = append(attlist, attempts)
	wordlist = append(wordlist, word)
	EncodSB(userlist, attlist, wordlist)
	return userlist, attlist, wordlist
}

func Result(userlist []string, attlist []int, wordlist []string) []string {
	var score []string 
	for i := 0; i < len(userlist); i++ {
		score = append(score, userlist[i] +  "			" + strconv.Itoa(attlist[i]) + "			" + wordlist[i])
	}
	return score
}

func ChooseImage(attempts int) string {
	img := ""
	switch attempts {
	case 10: 
		img = "/static/hangmanimg/0.jpg"
		
	case 9:
		img ="/static/hangmanimg/1.jpg"
		
	case 8:
		img = "/static/hangmanimg/2.jpg"
		
	case 7:
		img = "/static/hangmanimg/3.jpg"
		
	case 6:
		img = "/static/hangmanimg/4.jpg"
		
	case 5:
		img = "/static/hangmanimg/5.jpg"
		
	case 4:
		img ="/static/hangmanimg/6.jpg"
		
	case 3:
		img = "/static/hangmanimg/7.jpg"
		
	case 2:
		img = "/static/hangmanimg/8.jpg"
		
	case 1:
		img = "/static/hangmanimg/9.jpg"
		
	default:
		img = "/static/hangmanimg/10.jpg"
		
	}
	return img
}


