package piscine

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
