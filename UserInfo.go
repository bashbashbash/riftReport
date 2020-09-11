package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiKey string = "RGAPI-6860c312-4a43-4556-b7ae-478a0291b64a"

// UserInfo resonse body json
// API SUMMONER-V4
type UserInfo struct {
	ID            string
	AccountID     string
	Puuid         string
	Name          string
	ProfileIconID int
	RevisionDate  int
	SummonerLevel int
}

// league entries in all queues for a given summoner ID response body json
// LEAGUE-V4
type LeagueEntries struct {
	LeagueID     string
	SummonerID   string
	SummonerName string
	QueueType    string
	Tier         string
	Rank         string
	LeaguePoints int
	Wins         int
	Losses       int
	HotStreak    bool
	Veteran      bool
	Freshblood   bool
	Inactive     bool
}

func main() {
	userInfo := summonerInfo("bashbashbash")

	userInfoJSON, _ := json.MarshalIndent(userInfo, "", "  ")
	fmt.Printf("MarshalIndent funnction output\n %s\n", string(userInfoJSON))

	details := summonerDetails(userInfo.ID)
	detailsJSON, _ := json.MarshalIndent(details, "", "  ")
	fmt.Printf("MarshalIndent funnction output\n %s\n", string(detailsJSON))
}

func summonerDetails(accountID string) LeagueEntries {
	resp, err := http.Get("https://na1.api.riotgames.com/lol/league/v4/entries/by-summoner/1xcRetZtGjvYhsN45YxKY4jxRWxLvFXywfEJ7Vm0VzfFjBc?api_key=RGAPI-6860c312-4a43-4556-b7ae-478a0291b64a")
	printIferr(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	printIferr(err)

	var leagueEntries []LeagueEntries
	err = json.Unmarshal(body, &leagueEntries)
	printIferr(err)

	return leagueEntries[0]
}

func summonerInfo(summonerName string) UserInfo {
	resp, err := http.Get("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-account/PFo90JxR0-a7fttW-cSYNSoJQZ15uNQp-W44WlCMELwxeHA?api_key=RGAPI-6860c312-4a43-4556-b7ae-478a0291b64a")
	printIferr(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	printIferr(err)

	var userInfo UserInfo
	err = json.Unmarshal(body, &userInfo)
	printIferr(err)

	return userInfo
}

func printIferr(err error) {
	if err != nil {
		print(err)
	}
}
