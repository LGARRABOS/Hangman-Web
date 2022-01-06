package piscine

import (
	"os"
	"encoding/json"
	"log"

)

type Save struct {
    Attempts  int `json:"attempts"`
	Word      string `json:"word"`
	Stock     string `json:"stock"`
	Asciiart  string `json:"asciiart"`
}

func Encod(a int, w string, s []byte, ascii string) {
	use := ""
	for i:= 0; i < len(s); i++ {
		use = use + string(s[i])
	}
	savegame := Save{Attempts: a, Word: w, Stock: use, Asciiart: ascii}
	json_data, err := json.Marshal(savegame)
	err2 := os.WriteFile("save.txt", json_data, 0666)
    if err != nil {
    	log.Fatal(err)
    }
	if err2 != nil {
    	log.Fatal(err2)
    }
}

func Decod() (int, string, []byte, string) {
	var restore Save
	var stock []byte
	json_data, err := os.ReadFile(os.Args[2])
	if err != nil {
    	log.Fatal(err)
    }
	err2 := json.Unmarshal(json_data, &restore)
	if err2 != nil {
    	log.Fatal(err2)
    }
	for i := 0; i < len(restore.Stock); i++ {
		stock = append(stock, restore.Stock[i])
	}
	return restore.Attempts, restore.Word, stock, restore.Asciiart


}