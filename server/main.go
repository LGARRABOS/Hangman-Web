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

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	

	http.HandleFunc("/hangman", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("../index.html"))
	letter := ""
	var hidden_word string
	var attempts int
	var win bool
	won := ""
	start := true
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
		letter = l[0]	
		if start {
			hidden_word, attempts, win = hangman.Hangman(letter, true)
			start = false
		}
	}
	if win {
		won = "Congrats !"
	}
	data := Data {
		LetterChoose: letter,
		Attempts: attempts,
		TabUnderscore: hidden_word,
		Won: won,
	}
	fmt.Println(hidden_word, attempts, win)
	tmpl.Execute(w, data)
}