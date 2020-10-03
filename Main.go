package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// contains data on all champions
var champions map[string]interface{}

func main() {
	// position 1 argument should be summoner name
	// TODO add validation
	process(os.Args[1])

}

func process(name string) {
	userInfo := getUserInfo(name)
	var jsonObject interface{}
	jsonObject = getJSONDataFromResp(userInfo, jsonObject)
	printStructObject(jsonObject)
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

// get game
