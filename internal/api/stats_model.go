package api

type PlayerStatsResponse struct {
	PlayerID string    `json:"player_id"`
	GameID   string    `json:"game_id"`
	Lifetime Lifetime  `json:"lifetime"`
	Segments []Segment `json:"segments"`
}

type Lifetime struct {
	LongestWinStreak    string   `json:"Longest Win Streak"`
	Wins                string   `json:"Wins"`
	AverageKDRatio      string   `json:"Average K/D Ratio"`
	WinRate             string   `json:"Win Rate %"`
	CurrentWinStreak    string   `json:"Current Win Streak"`
	RecentResults       []string `json:"Recent Results"`
	TotalHeadshotsPct   string   `json:"Total Headshots %"`
	Matches             string   `json:"Matches"`
	KDRatio             string   `json:"K/D Ratio"`
	AverageHeadshotsPct string   `json:"Average Headshots %"`
}

type Segment struct {
	ImgRegular string `json:"img_regular"`
	Stats      Stats  `json:"stats"`
	Type       string `json:"type"`
	Mode       string `json:"mode"`
	Label      string `json:"label"`
	ImgSmall   string `json:"img_small"`
}

type Stats struct {
	AverageKDRatio      string `json:"Average K/D Ratio"`
	AverageTripleKills  string `json:"Average Triple Kills"`
	AverageQuadroKills  string `json:"Average Quadro Kills"`
	AverageDeaths       string `json:"Average Deaths"`
	AverageHeadshotsPct string `json:"Average Headshots %"`
	AverageMVPs         string `json:"Average MVPs"`
	Assists             string `json:"Assists"`
	AverageKRRatio      string `json:"Average K/R Ratio"`
	WinRate             string `json:"Win Rate %"`
	Wins                string `json:"Wins"`
	AverageAssists      string `json:"Average Assists"`
	QuadroKills         string `json:"Quadro Kills"`
	TripleKills         string `json:"Triple Kills"`
	MVPs                string `json:"MVPs"`
	KDRatio             string `json:"K/D Ratio"`
	HeadshotsPerMatch   string `json:"Headshots per Match"`
	Rounds              string `json:"Rounds"`
	PentaKills          string `json:"Penta Kills"`
	Matches             string `json:"Matches"`
	Kills               string `json:"Kills"`
	Headshots           string `json:"Headshots"`
	Deaths              string `json:"Deaths"`
	AveragePentaKills   string `json:"Average Penta Kills"`
	KRRatio             string `json:"K/R Ratio"`
	AverageKills        string `json:"Average Kills"`
	TotalHeadshotsPct   string `json:"Total Headshots %"`
}
