package main

import (
	"fmt"
	h "hangman/hangman"
	"html/template"
	"net/http"
	piscine "hangman/function"
)

type Data struct {
	LetterChoose  string
	Attempts      int
	TabUnderscore string
	Won           string
	Used string
	WordToFind string
	UsrList []string
	Image string
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/hangman", Handler)

	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe("localhost:8080", nil)

}

var (
	start       = 0
	letter      = ""
	hidden_word string
	won         = ""
	word 		= ""
	user string
	findword string
	userlist []string
	attlist []int
	wordlist []string
	difflist []string
	result []string
	img string
	imgpath = false
	usr []string
	home = template.Must(template.ParseFiles("../home.html"))
	hangman = template.Must(template.ParseFiles("../index.html"))
	accueil = template.Must(template.ParseFiles("../accueil.html"))
	end = template.Must(template.ParseFiles("../final.html"))
	score = template.Must(template.ParseFiles("../scoreboard.html"))
	diff string
)

func Handler(w http.ResponseWriter, r *http.Request) {


	
	win := false
	letter := ""
	stock := ""
	attempts := 10

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
				start = 1
				usr = m["username"]
				if len(usr) != 0 {
					user = usr[0]
				}
			} else if redir[0] == "scoreboard" {
				start = 4
			}
		} 
		if len(level) !=  0  {
			if level[0] == "easy" {
				word = "words.txt"
				diff = "easy"
			} else if level[0] == "medium" {
				word = "words2.txt"
				diff = "medium"
			} else {
				word = "words3.txt"
				diff = "hard"
			}
			hidden_word, attempts, win, won, stock, findword = h.Hangman(letter, true, word, user, diff)
			start = 2
			img = piscine.ChooseImage(attempts)
			imgpath = true
		}
		if len(l) > 0 {
			letter = l[0]
			hidden_word, attempts, win, won, stock, findword = h.Hangman(letter, false, word, user, diff)
			img = piscine.ChooseImage(attempts)
			imgpath = true
		}
		
		if win {
			won = "win"
			start = 3
			userlist, attlist, wordlist, difflist= piscine.AddScore(user, attempts, findword, diff)
			result = piscine.Result(userlist, attlist, wordlist, difflist)
		}
		if attempts <= 0 {
			won = "loose"
			start = 3
			userlist, attlist, wordlist, difflist = piscine.AddScore(user, attempts, findword, diff)
			fmt.Println(userlist, attlist, wordlist)
			result = piscine.Result(userlist, attlist, wordlist, difflist)
		}
	}
	data := Data{
		LetterChoose:  letter,
		Attempts:      attempts,
		TabUnderscore: hidden_word,
		Won:           won,
		Used: stock,
		WordToFind: findword,
		UsrList: result,
		Image: img,
	}
	if start == 0 {
		accueil.Execute(w, data)
	} else if start == 1 {
		home.Execute(w, data)
	} else if start == 2 {
		hangman.Execute(w, data)
	} else if start == 3 {
		end.Execute(w, data)
	} else if start  == 4 {
		score.Execute(w, data)
	}
	
}