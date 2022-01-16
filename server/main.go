package main

import (
	"fmt"
	h "hangman/hangman"
	"html/template"
	"net/http"
)

type Data struct {
	LetterChoose  string
	Attempts      int
	TabUnderscore string
	Won           string
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/hangman", Handler)

	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)

}

var (
	start       = 0
	letter      = ""
	hidden_word string
	attempts    int
	win         bool
	won         = ""
	word 		= ""
)

func Handler(w http.ResponseWriter, r *http.Request) {

	home := template.Must(template.ParseFiles("../home.html"))
	hangman := template.Must(template.ParseFiles("../index.html"))
	accueil := template.Must(template.ParseFiles("../accueil.html"))
	end := template.Must(template.ParseFiles("../final.html"))

	switch r.Method {
	case "GET":
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Printf("POST METHOD = %v\n", r.PostForm)
		m := r.PostForm
		l := m["letter"]
		level := m["level"]
		redir := m ["redir"]
		if len(redir) != 0 {
			if redir[0] == "accueil" {
				start = 0
			} else if redir[0] == "home" {
				fmt.Println("85")
				
				start = 1
				fmt.Println("81")
			}
		} 
		if len(level) !=  0 {
			if level[0] == "easy" {
				word = "words.txt"
			} else if level[0] == "medium" {
				word = "words2.txt"
			} else {
				word = "words3.txt"
			}
			hidden_word, attempts, win, won = h.Hangman(letter, true, word)
			start = 2
		}
		if len(l) > 0 {
			letter = l[0]
		}
		hidden_word, attempts, win, won = h.Hangman(letter, false, word)
		
		if win {
			won = "Congrats !"
			start = 3
		}
		if attempts <= 0 {
			won = "The poor José is dead because of you !"
			start = 3
		}
	}
	data := Data{
		LetterChoose:  letter,
		Attempts:      attempts,
		TabUnderscore: hidden_word,
		Won:           won,
	}
	letter = ""
	if start == 0 {
		accueil.Execute(w, data)
	} else if start == 1 {
		home.Execute(w, data)
	} else if start == 2 {
		hangman.Execute(w, data)
	} else if start == 3 {
		end.Execute(w, data)
	}
}
