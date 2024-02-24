package config

import (
	"os"

	"github.com/Santiago-Balcero/go-spotify/utils"
	"github.com/joho/godotenv"
)

var SpotifyId string
var SpotifySecret string

func LoadConfig() {
	err := godotenv.Load(".env")
	utils.CheckError(err)
	SpotifyId = os.Getenv("SPOTIFY_ID")
	SpotifySecret = os.Getenv("SPOTIFY_KEY")
}
