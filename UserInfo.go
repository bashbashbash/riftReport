package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiKey string = "RGAPI-738ddbf9-45a4-452c-b26b-2bb6ee8fcade"

// UserInfo resonse body json
type UserInfo struct {
	ID            string
	AccountID     string
	Puuid         string
	Name          string
	ProfileIconID int
	RevisionDate  int
	SummerLevel   int
}

func main() {
	userInfo := summonerInfo("bashbashbash")

	fmt.Printf("%s is the name of the user and the PUUID is: %s", userInfo.Name, userInfo.Puuid)
}

func summonerInfo(summonerName string) UserInfo {
	resp, err := http.Get("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/" + summonerName + "?api_key=" + apiKey)
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
