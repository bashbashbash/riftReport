package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// contains data on all champions
var champions map[string]interface{}
var encryptedSummonerID string
var encryptedAccountID string

func main() {
	validateApiKey()
	// position 1 argument should be summoner name
	// TODO add validation
	process(os.Args[1])
}

func validateApiKey() {

}

func process(summonerID string) {
	getEncryptedKeys(summonerID)
}
func getEncryptedKeys(summonerID string) {
	userInfo := getUserInfo(summonerID)
	var jsonObject interface{}
	jsonObject = getJSONDataFromResp(userInfo, jsonObject)
	setEncryptedSummonerID(getValueFromJSONObject(jsonObject, "id"))
	setEncryptedAccountID(getValueFromJSONObject(jsonObject, "accountId"))
}

// printUserInfo - using a summonerID, print UserInfo Struct which includes
// Encrypted accountID, Encrypted summonerID, summonerID, ProfileIconID, PUUID, RevisionDate, and SummonerLevel
func printUserInfo(summonerID string) {
	userInfo := getUserInfo(summonerID)
	var jsonObject interface{}
	jsonObject = getJSONDataFromResp(userInfo, jsonObject)
	printStructObject(jsonObject)
}

func getValueFromJSONObject(jsonObject interface{}, key string) string {
	casting := jsonObject.(map[string]interface{})
	return casting[key].(string)
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
