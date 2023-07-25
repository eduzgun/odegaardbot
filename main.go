/*
package main

import (
	"fmt"
	"math/rand"
	"regexp"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	//https://optaplayerstats.statsperform.com/en_GB/soccer/premier-league-2022-2023/80foo89mm28qjvyhjzlpwj28k/match/view/7t5r2nwlkdj3c1vrc5ph22jo4/match-summary
	//quotes := []Quote{}
	escaped := regexp.QuoteMeta("sofascore.com/arsenal-manchester-united/KR")
	r := regexp.MustCompile(`^https?:\/\/[a-z]*\.?` + escaped + `.*`)
	c := colly.NewCollector(
		colly.URLFilters(r),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)

	extensions.RandomMobileUserAgent(c)
	extensions.Referer(c)

	c.OnRequest(func(r *colly.Request) {
		// put a forward slash and apostrophe here r.Headers.Set("Accept", "*")

		fmt.Println("visiting url", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {

		fmt.Println("accessed website")
		fmt.Println("Response code: ", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML("span.sc-bqWxrE.jsASkT", func(h *colly.HTMLElement) {
		fmt.Println("accessed website")

		fmt.Println(h.Text)

	})

	fmt.Println("started")
	err := c.Visit("https://www.sofascore.com/arsenal-manchester-united/KR")
	if err != nil {
		fmt.Printf("failed to visit url: %v\n", err)
		return
	}
}

*/

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

	passAccuracy100 := (ode.Statistics.AccuratePass / ode.Statistics.TotalPass)
	fmt.Println("Pass accuracy: ", passAccuracy100)
	fmt.Println("Error: ", err2)
	fmt.Println("XG: ", ode.Statistics.XG)
}
