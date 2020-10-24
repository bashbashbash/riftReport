package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/gommon/log"
)

// contains data on all champions
var champions map[string]interface{}
var encryptedSummonerID string
var encryptedAccountID string

func main() {
	// position 1 argument should be summoner name
	// TODO add validation
	process(os.Args[1])
}

func process(summonerID string) {
	getEncryptedKeys(summonerID)
	printGameHistory()
	printWins(summonerID)
}

func printWins(summonerID string) {
	gamesWon := countLastNWins(10)
	fmt.Printf("Player %v has one the past %d out of %d games", summonerID, gamesWon, 10)
}

func countLastNWins(numberOfMatchesToReview int) int {
	return numberOfMatchesToReview
}
func printGameHistory() {
	getMatches()
}

func getMatches() {
	matchHistory := getMatchHistory()
	var jsonObject interface{}
	matchListDtoJSONObject, err := getJSONDataFromResp(matchHistory, jsonObject)
	if err != nil {
		log.Error("failed to populate userInfoJSON")
	}
	printStructObject(matchListDtoJSONObject)
}
func getEncryptedKeys(summonerID string) {
	userInfo := getUserInfo(summonerID)
	var jsonObject interface{}
	jsonObject, err := getJSONDataFromResp(userInfo, jsonObject)
	if err != nil {
		log.Error("failed to populate userInfoJSON")

	}
	setEncryptedSummonerID(getValueFromJSONObject(jsonObject, "id"))
	setEncryptedAccountID(getValueFromJSONObject(jsonObject, "accountId"))
}

// printUserInfo - using a summonerID, print UserInfo Struct which includes
// Encrypted accountID, Encrypted summonerID, summonerID, ProfileIconID, PUUID, RevisionDate, and SummonerLevel
func printUserInfo(summonerID string) {
	userInfo := getUserInfo(summonerID)
	var jsonObject interface{}
	userInfoJSON, err := getJSONDataFromResp(userInfo, jsonObject)
	if err != nil {
		log.Error("failed to populate userInfoJSON")
	}
	printStructObject(userInfoJSON)
}

func getValueFromJSONObject(jsonObject interface{}, key string) string {
	casting := jsonObject.(map[string]interface{})
	return casting[key].(string)
}

func getJSONDataFromResp(requestURL string, jsonContainer interface{}) (interface{}, error) {
	resp, err := http.Get(requestURL)
	printIferr(err)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	printIferr(err)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &jsonContainer)
	printIferr(err)
	if err != nil {
		return nil, err
	}

	return jsonContainer, nil
}

func getEncryptedSummonerID() string {
	return encryptedSummonerID
}

func getEncryptedAccountID() string {
	return encryptedAccountID
}

func setEncryptedAccountID(eAccountID string) {
	encryptedAccountID = eAccountID
}

func setEncryptedSummonerID(eSummonerID string) {
	encryptedSummonerID = eSummonerID
}
