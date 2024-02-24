package models

type Artist struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Popularity int      `json:"popularity"`
	Albums     []Album  `json:"albums"`
	Genres     []string `json:"genres"`
	Url        string   `json:"url"`
	Followers  int      `json:"followers"`
	Image      string   `json:"image"`
}

type Album struct {
	Name   string  `json:"name"`
	Tracks []Track `json:"tracks"`
}

type Track struct {
	Name         string  `json:"name"`
	Danceability float32 `json:"danceability"`
	Energy       float32 `json:"energy"`
}
