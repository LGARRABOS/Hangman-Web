package piscine

import (
	"strconv"
)
/*
Transforme les tableaux de byte en sstring,
sert pour transformer le stock des lettres en string pour le faire apparaitre dans le site.
*/
func TabToString(stock []byte) string {
	strstock := ""
	for i := 0; i < len(stock); i++ {
		strstock = strstock + string(stock[i]) + ", "
	}
	return strstock
}
/*
Ajoute les nouvelles données dans le JSON pour sotcker les score.
*/
func AddScore(user string, attempts int, word string, diff  string) ([]string, []int, []string, []string) {
	userlist, attlist, wordlist, difflist:= DecodSB()
	userlist = append(userlist, user)
	attlist = append(attlist, attempts)
	wordlist = append(wordlist, word)
	difflist = append(difflist, diff)
	EncodSB(userlist, attlist, wordlist, difflist)
	return userlist, attlist, wordlist, difflist
}
/*
transforme le scoreboard en un tableau de string pour print 
ligne par lignesur le site web.
*/
func Result(userlist []string, attlist []int, wordlist []string, difflist []string) []string {
	var score []string 
	for i := 0; i < len(userlist); i++ {
		score = append(score, userlist[i] +  " " + strconv.Itoa(attlist[i]) + " " + wordlist[i] + " " + difflist[i])
	}
	return score
}
/* 
choisir l'image qui sera afficher sur index.html
selon le nombre de chance restante
*/
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


