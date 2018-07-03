package ov

type FootballInfo struct {
	Id       string
	Name     string
	Selected string
	Info     []FootballNatchDateInfo
}
type FootballNatchDateInfo struct {
	Date string
	Info []FootballNatchInfo
}
type FootballNatchInfo struct {
	Id             string
	Date           string
	StartTime      string
	HomeName       string
	HomeLogo       string
	SeasonId       string
	RoundId        string
	HasProcessLogo string
	GuestName      string
	GuestLogo      string
	GuestGoal      string
	LeagueId       string
}
