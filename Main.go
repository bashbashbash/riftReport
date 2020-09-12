package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	summonerName := "bashbashbash"
	var emptyUserInfo UserInfo
	userInfo := getJSONDataFromResp("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/"+summonerName+"?api_key="+apiKey, emptyUserInfo)

	userInfoJSON, _ := json.MarshalIndent(&userInfo, "", "  ")
	fmt.Printf("MarshalIndent funnction output\n %s\n", string(userInfoJSON))

	// details := summonerDetails(userInfo.ID)
	// detailsJSON, _ := json.MarshalIndent(details, "", "  ")
	// fmt.Printf("MarshalIndent funnction output\n %s\n", string(detailsJSON))
}

func summonerDetails(accountID string) LeagueEntries {
	resp, err := http.Get()
	printIferr(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	printIferr(err)

	var leagueEntries []LeagueEntries
	err = json.Unmarshal(body, &leagueEntries)
	printIferr(err)

	return leagueEntries[0]
}

func getJSONDataFromResp(requestURL string, jsonContainer interface{}) interface{} {
	resp, err := http.Get(requestURL)
	printIferr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	printIferr(err)

	err = json.Unmarshal(body, &jsonContainer)
	printIferr(err)
	return jsonContainer
}
