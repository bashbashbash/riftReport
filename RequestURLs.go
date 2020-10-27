package main

import (
	"strconv"
	"strings"
)

// USERINFO - get request
const USERINFO = "https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/***?api_key=***"

// OVERALLSTATS - get request
const OVERALLSTATS = "https://na1.api.riotgames.com/lol/league/v4/entries/by-summoner/***?api_key=***"

// ALLCHAMPDATA - get request
const ALLCHAMPDATA = "http://ddragon.leagueoflegends.com/cdn/10.18.1/data/en_US/champion.json"

// MATCHHISTORY - get request
const MATCHHISTORY = "https://na1.api.riotgames.com/lol/match/v4/matchlists/by-account/***?endIndex=***&api_key=***"

// SINGLEMATCHRECORD - get request
const SINGLEMATCHRECORD = "https://na1.api.riotgames.com/lol/match/v4/matches/***?api_key=***"

func getUserInfo(summonerName string) string {
	userInfoWithSummonerName := strings.Replace(USERINFO, "***", summonerName, 1)
	userInfoWithSummonerName = strings.Replace(userInfoWithSummonerName, "***", apiKey, 1)
	return userInfoWithSummonerName
}

func getMatchHistory() string {
	reqMatHist := strings.Replace(MATCHHISTORY, "***", getEncryptedAccountID(), 1)
	reqMatHist = strings.Replace(reqMatHist, "***", "10", 1)
	reqMatHist = strings.Replace(reqMatHist, "***", apiKey, 1)
	return reqMatHist
}

func getSingleMatchHistory(matchID int) string {
	matchRecord := strings.Replace(SINGLEMATCHRECORD, "***", strconv.Itoa(matchID), 1)
	matchRecord = strings.Replace(matchRecord, "***", apiKey, 1)
	return matchRecord
}
