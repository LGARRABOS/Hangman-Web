package piscine

func Glue(finalword string, attempts int, win bool) string {
	word := ""
	if attempts <= 0 || win == true {
		for i := 0; i < len(finalword); i ++ {
			if finalword[i] < 123 && finalword[i] > 96 {
				word = word + string(finalword[i])
			}
		}
		return word
	}
	return finalword	
}

func TabToString(stock []byte) string {
	strstock := ""
	for i := 0; i < len(stock); i++ {
		strstock = strstock + string(stock[i]) + ", "
	}
	return strstock
}