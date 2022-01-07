package main

import (
		/*"html/template"
        "log"
        "net/http"
		"fmt"*/
		hangman "hangman/hangman-classic/hangman"
)
func main() {
        /*// Set routing rules
        http.HandleFunc("/", Tmp)

        //Use the default DefaultServeMux.
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
                log.Fatal(err)
        }*/
		hangman.Hangman();

}
/*
func Tmp(w http.ResponseWriter, r *http.Request) { 
	tmpl, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tmpl.Execute(w , nil)
}*/