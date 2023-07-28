package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"strconv"
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

	req.Header.Add("X-RapidAPI-Key", "76a5c68254msh15c95aa5f37d156p1a408ajsn7bc8944039ed")
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

	fmt.Println("Accurate passes: ", ode.Statistics.AccuratePass)
	fmt.Println("Total passes: ", ode.Statistics.TotalPass)

	passAccuracy100 := (float32(ode.Statistics.AccuratePass) / float32(ode.Statistics.TotalPass))
	fmt.Println("Pass accuracy: ", passAccuracy100)
	fmt.Println("Error: ", err2)
	fmt.Println("XG: ", ode.Statistics.XG)
}
