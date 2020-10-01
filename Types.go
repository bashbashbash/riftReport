package main

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

// LeagueEntries in all queues for a given summoner ID response body json
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

// ChampionStatsData json container for champion stats
type ChampionStatsData struct {
	Hp                   int
	Hpperlevel           int
	Mp                   int
	Mpperlevel           int
	Movespeed            int
	Armor                int
	Armorperlevel        float64
	Spellblock           float64
	Spellblockperlevel   float64
	Attackrange          int
	Hpregen              int
	Hpregenperlevel      int
	Mpregen              int
	Mpregenperlevel      int
	Crit                 int
	Critperlevel         int
	Attackdamage         int
	Attackdamageperlevel int
	Attackspeedperlevel  float64
	Attackspeed          float64
}

// ChampionTagData json container for role types
type ChampionTagData struct {
	Tags []string
}

// ChampionImageData json container for champion image data
type ChampionImageData struct {
	Full   string
	Sprite string
	Group  string
	X      int
	Y      int
	W      int
	H      int
}

// ChampionInfoData json container for champion base stats
type ChampionInfoData struct {
	Attack     int
	Defense    int
	Magic      int
	Difficulty int
}

// ChampionData json container for champion data
type ChampionData struct {
	Version string
	ID      string
	Key     string
	Name    string
	Title   string
	Blurb   string
	Info    ChampionInfoData
	Image   ChampionImageData
	Tag     ChampionTagData
	ParType string
	Stats   ChampionStatsData
}

// ChampionName json container for each champion
type ChampionName struct {
	Name ChampionData
}

// Champions top level json container
type Champions struct {
	Type    string
	Format  string
	Version string
	Data    ChampionName
}

// MatchlistDto - https://developer.riotgames.com/apis#match-v4/GET_getMatchlist
type MatchlistDto struct {
	startIndex int
	totalGames int
	endIndex   int
	matches    MatchReferenceDto
}

// MatchReferenceDto - https://developer.riotgames.com/apis#match-v4/GET_getMatchlist
type MatchReferenceDto struct {
	gameId     int64
	role       string
	season     int
	platformId string
	champion   int
	queue      int
	lane       string
	timestamp  int64
}