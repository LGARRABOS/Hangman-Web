package main

import (
		"html/template"
        "net/http"
		"fmt"
		hangman "hangman/hangman"
)
type Data struct {
	LetterChoose string
	Attempts int
	TabUnderscore string
	Won string

}

func main() {
	http.HandleFunc("/", Handler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	http.HandleFunc("/hangman", Handler)
	http.ListenAndServe(":8080", nil)

}
var (
	start bool = true
	c int = 0
	letter = ""
	hidden_word string
	attempts int
	win bool
	won = ""
)

func Handler(w http.ResponseWriter, r *http.Request) {
	

	tmpl := template.Must(template.ParseFiles("../index.html"))
	
	switch r.Method {
	case "GET" :
		fmt.Println("GET METHOD")
	case "POST" :
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Printf("POST METHOD = %v\n", r.PostForm)
		m := r.PostForm
		l := m["letter"]
		if len(l) > 0 {
			letter = l[0]
		}
		if start {
			fmt.Println(letter)
			hidden_word, attempts, win = hangman.Hangman(letter, true )
			start = false
		} else {
			fmt.Println(letter)
			hidden_word, attempts, win = hangman.Hangman(letter, false)
		} 
		
	}
	c++
	if win {
		won = "Congrats !"
		start = true
	}
	data := Data {
		LetterChoose: letter,
		Attempts: attempts,
		TabUnderscore: hidden_word,
		Won: won,
	}
	fmt.Println(hidden_word, attempts, win, c)
	tmpl.Execute(w, data)
}