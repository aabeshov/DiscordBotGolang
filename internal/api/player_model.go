package api

import (
	"time"
)

type PlayerReponse struct {
	PlayerID           string            `json:"player_id"`
	Nickname           string            `json:"nickname"`
	Avatar             string            `json:"avatar"`
	Country            string            `json:"country"`
	CoverImage         string            `json:"cover_image"`
	Platforms          Platforms         `json:"platforms"`
	Games              map[string]Game   `json:"games"`
	Settings           Settings          `json:"settings"`
	FriendsIDs         []string          `json:"friends_ids"`
	NewSteamID         string            `json:"new_steam_id"`
	SteamID64          string            `json:"steam_id_64"`
	SteamNickname      string            `json:"steam_nickname"`
	Memberships        []string          `json:"memberships"`
	FaceitURL          string            `json:"faceit_url"`
	MembershipType     string            `json:"membership_type"`
	CoverFeaturedImage string            `json:"cover_featured_image"`
	Infractions        map[string]string `json:"infractions"` // Assuming infractions is a map, adjust if necessary
	Verified           bool              `json:"verified"`
	ActivatedAt        time.Time         `json:"activated_at"`
}

type Platforms struct {
	Steam string `json:"steam"`
}

type Game struct {
	Region          string                 `json:"region"`
	GamePlayerID    string                 `json:"game_player_id"`
	SkillLevel      int                    `json:"skill_level"`
	FaceitElo       int                    `json:"faceit_elo"`
	GamePlayerName  string                 `json:"game_player_name"`
	SkillLevelLabel string                 `json:"skill_level_label"`
	Regions         map[string]interface{} `json:"regions"` // Adjust if the structure of regions is known
	GameProfileID   string                 `json:"game_profile_id"`
}

type Settings struct {
	Language string `json:"language"`
}
