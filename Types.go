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
	gameID     int64
	role       string
	season     int
	platformID string
	champion   int
	queue      int
	lane       string
	timestamp  int64
}

// MatchDto - Get match by match ID - https://developer.riotgames.com/apis#match-v4/GET_getMatch
type MatchDto struct {
	gameID                int64
	participantIdentities []ParticipantIdentityDto
	queueID               int
	gameType              string
	gameDuration          int64
	teams                 []TeamStatsDto
	platformID            string
	gameCreation          int64
	seasonID              int
	gameVersion           string
	mapID                 int
	gameMode              string
	participants          []ParticipantDto
}

// ParticipantIdentityDto - see MatchDto
type ParticipantIdentityDto struct {
	partiticpantID int
	player         PlayerDto
}

// PlayerDto - see ParticipantIdentityDto
type PlayerDto struct {
	profileIcon       int
	accountID         string
	matchHistoryURI   string
	currentAccountID  string
	currentPlatformID string
	summonerName      string
	summonerID        string
	platformID        string
}

// TeamStatsDto - see MatchDto
type TeamStatsDto struct {
	towerKills           int
	riftHeraldKills      int
	firstBlood           bool
	inhibitorKills       int
	bans                 []TeamBansDto
	firstBaron           bool
	firstDragon          bool
	dominionVictoryScore int
	dragonKills          int
	baronKills           int
	firstInhibitor       bool
	firstTower           bool
	vilemawKills         int
	firstRiftHerald      bool
	teamID               int
	win                  string
}

// TeamBansDto - see PlayerDto
type TeamBansDto struct {
	championID int
	pickTurn   int
}

// ParticipantDto - see MatchDto
type ParticipantDto struct {
	participantID             int
	championID                int
	runes                     []RuneDto
	stats                     ParticipantStatsDto
	teamID                    int
	timeline                  ParticipantTimelineDto
	spell1ID                  int
	spell2ID                  int
	highestAchievedSeasonTier string
	masteries                 []MasteryDto
}

// RuneDto - see ParticipantDto
type RuneDto struct {
	runeID int
	rank   int
}

// ParticipantStatsDto - see ParticipantDto
type ParticipantStatsDto struct {
	item0                           int
	item2                           int
	totalUnitsHealted               int
	item1                           int
	largestMultiKill                int
	goldEarned                      int
	firstInhibitorKill              bool
	physicalDamageTaken             int64
	nodeNeutralizeAssist            int
	totalPlayerScore                int
	champLevel                      int
	damageDealtToObjectives         int64
	totalDamageTaken                int64
	neutralMinionsKilled            int
	deaths                          int
	tripleKills                     int
	magicDamageDealtToChampions     int64
	wardsKilled                     int
	pentaKills                      int
	damageSelfMitigated             int64
	largestCriticalStrike           int
	nodeNeutralize                  int
	totalTimeCrowdControlDealt      int
	firstTowerKill                  bool
	magicDamageDealt                int64
	totalScoreRank                  int
	nodeCapture                     int
	wardsPlaced                     int
	totalDamageDealt                int64
	timeCCingOther                  int64
	magicalDamageTaken              int64
	largestKillingSpree             int
	totalDamageDealtToChampions     int64
	physicalDamaDealtToChampions    int64
	neutralMinionsKilledTeamJungle  int
	totalMinionsKilled              int
	firstInhibitorAssist            bool
	visionWardsBoughtInGame         int
	objectivePlayerScore            int
	kills                           int
	firstTowerAssist                bool
	combatPlayerScore               int
	inhibitorKills                  int
	turrentKills                    int
	participantID                   int
	trueDamageTaken                 int64
	firstBloodAssist                bool
	nodeCaptureAssist               int
	assists                         int
	teamObjective                   int
	altarsNeutralized               int
	goldSpent                       int
	damageDealtToTurrents           int64
	altarsCaptured                  int
	win                             bool
	totalHeal                       int64
	unrealKills                     int
	visionScore                     int64
	physicalDamageDealt             int64
	firstBloodKill                  bool
	longestTimeSpentLiving          int
	killingSprees                   int
	sightWardsBoughtInGame          int
	trueDamageDealtToChampions      int64
	neutralMinionsKilledEnemyJungle int
	doubleKills                     int
	trueDamageDealt                 int64
	quadraKills                     int
	item4                           int
	item3                           int
	item6                           int
	item5                           int
	playerScore0                    int
	playerScore1                    int
	playerScore2                    int
	playerScore3                    int
	playerScore4                    int
	playerScore5                    int
	playerScore6                    int
	playerScore7                    int
	playerScore8                    int
	playerScore9                    int
	perk0                           int
	perk0Var1                       int
	perk0Var2                       int
	perk0Var3                       int
	perk1                           int
	perk1Var1                       int
	perk1Var2                       int
	perk1Var3                       int
	perk2                           int
	perk2Var1                       int
	perk2Var2                       int
	perk3                           int
	perk3Var1                       int
	perk3Var2                       int
	perk3Var3                       int
	perk4                           int
	perk4Var1                       int
	perk4Var2                       int
	perk4Var3                       int
	perk5                           int
	perk5Var1                       int
	perk5Var2                       int
	perk5Var3                       int
	perkPrimaryStyle                int
	perkSubStyle                    int
	statPerk1                       int
	statPerk2                       int
}

// ParticipantTimelineDto - see ParticipantDto
type ParticipantTimelineDto struct {
	participantID               int
	csDiffPerMinDeltas          map[string]float64
	damageTakenPerMinDeltas     map[string]float64
	role                        string
	damageTakenDiffPerMinDeltas map[string]float64
	xpPerMinDeltas              map[string]float64
	xpDiffPerMinDeltas          map[string]float64
	lane                        string
	creepsPerMinDeltas          map[string]float64
	goldPerMinDeltas            map[string]float64
}

// MasteryDto - see ParticipantDto
type MasteryDto struct {
	rank     int
	masterID int
}
