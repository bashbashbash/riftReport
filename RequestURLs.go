package main

import "strings"

// USERINFO - get request
const USERINFO = "https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/***?api_key=***"

// OVERALLSTATS - get request
const OVERALLSTATS = "https://na1.api.riotgames.com/lol/league/v4/entries/by-summoner/***?api_key=***"

// ALLCHAMPDATA - get request
const ALLCHAMPDATA = "http://ddragon.leagueoflegends.com/cdn/10.18.1/data/en_US/champion.json"

// MATCHHISTORY - get request
const MATCHHISTORY = "https://na1.api.riotgames.com/lol/match/v4/matchlists/by-account/***?endIndex=***api_key=***"

func getRequests(summonerName string) []string {
	userInfoWithSummonerName := strings.Replace(USERINFO, "***", summonerName, 1)
	requests := []string{userInfoWithSummonerName, OVERALLSTATS, ALLCHAMPDATA, MATCHHISTORY}

	for _, r := range requests {
		replaceAt(len(requests)-1, r, apiKey)
	}

	return requests
}
