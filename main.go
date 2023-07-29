package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	//"strconv"
	"github.com/joho/godotenv"
)

type Odegaard struct {
	Statistics struct {
		Goals                  int     `json:"goals"`
		Rating                 float32 `json:"rating"`
		XG                     float32 `json:"expectedGoals"`
		Assist                 int     `json:"goalAssist"`
		AccuratePass           int     `json:"accuratePass"`
		TotalPass              int     `json:"totalPass"`
		PassAccuracy           float32
		Touches                int `json:"touches"`
		OnTargetScoringAttempt int `json:"OnTargetScoringAttempt"`
		OffTargetShots         int `json:"shotOffTarget"`
		BigChanceCreated       int `json:"bigChanceCreated"`
		DuelWon                int `json:"duelWon"`
		DuelLost               int `json:"duelLost"`
		ArialWon               int `json:"arialWon"`
		ArialLost              int `json:"arialLost"`
		BlockedScoringAttempt  int `json:"blockedScoringAttempt"`
		MinutesPlayed          int `json:"minutesPlayed"`
	} `json:"statistics"`
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	//10385450 - everton
	//10385451 - brentford arsenal
	//547410 - odegaard id
	//42 - arsenal id
	//premier league -17 cat 1
	//135666 - tossard id
	url := "https://footapi7.p.rapidapi.com/api/match/10385451/player/547410/statistics"
	//var topStoryID string = strconv.Itoa(topStories[0])
	req, _ := http.NewRequest("GET", url, nil)

	my_rapidAPIKey := goDotEnvVariable("my_rapidAPIkey")
	req.Header.Add("X-RapidAPI-Key", my_rapidAPIKey)
	req.Header.Add("X-RapidAPI-Host", "footapi7.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyStr := string(body)

	var ode Odegaard

	err2 := json.Unmarshal([]byte(bodyStr), &ode)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Displaying stats for Martin Odegaard:")
	fmt.Println("Accurate passes: ", ode.Statistics.AccuratePass)
	fmt.Println("Total passes: ", ode.Statistics.TotalPass)

	passAccuracy100 := (float32(ode.Statistics.AccuratePass) / float32(ode.Statistics.TotalPass))
	fmt.Println("Pass accuracy: ", passAccuracy100)

	fmt.Println("XG: ", ode.Statistics.XG)
}
