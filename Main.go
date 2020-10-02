package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// contains data on all champions
var champions map[string]interface{}

func main() {
	summonerName := os.Args
	var id string
	requests := getRequests(summonerName[1])
	for i := 0; i < 3; i++ {
		var jsonObject interface{}
		// can only update 1 if 0 has retrieved key
		if i == 1 {
			requests[i] = strings.ReplaceAll(requests[i], "***", id)
		}

		jsonObject = getJSONDataFromResp(requests[i], jsonObject)
		// store key for request[1]
		if i == 0 {
			id = jsonObject.(map[string]interface{})["id"].(string)
		}
		if i == 2 {
			champions = jsonObject.(map[string]interface{})
		}
		printStructObject(jsonObject)
	}
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
