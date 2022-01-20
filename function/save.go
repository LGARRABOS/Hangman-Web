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
}
type Scoreboard struct {
	Username	[]string `json:"username"`
	Attempts	[]int `json:"attempts"`
	Word		[]string`json:"word"`
	Difficulty	[]string`json:"difficulty"`
}

func Encod(a int, w string, s []byte) {
	
	use := ""
	for i:= 0; i < len(s); i++ {
		use = use + string(s[i])
	}
	savegame := Save{Attempts: a, Word: w, Stock: use}
	json_data, err := json.Marshal(savegame)
	err2 := os.WriteFile("save.txt", json_data, 0666)
    if err != nil {
    	log.Fatal(err)
    }
	if err2 != nil {
    	log.Fatal(err2)
    }
}

func Decod() (int, string, []byte) {
	var restore Save
	var stock []byte
	json_data, err := os.ReadFile("save.txt")
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
	return restore.Attempts, restore.Word, stock


}
func DecodSB() ([]string, []int, []string, []string) {
	
	var restore Scoreboard
	json_data, err := os.ReadFile("scoreboard.txt")
	
	if err != nil {
    	log.Fatal(err)
    }
	err2 := json.Unmarshal(json_data, &restore)
	
	if err2 != nil {
    	log.Fatal(err2)
    }
	return restore.Username, restore.Attempts, restore.Word, restore.Difficulty

}
func EncodSB(u []string, a []int, w []string, d []string) {
	savescore := Scoreboard{Username: u, Attempts: a, Word: w, Difficulty: d}
	json_data, err := json.Marshal(savescore)
	err2 := os.WriteFile("scoreboard.txt", json_data, 0666)
    if err != nil {
    	log.Fatal(err)
    }
	if err2 != nil {
    	log.Fatal(err2)
    }

} 