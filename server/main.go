package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Santiago-Balcero/go-spotify/config"
	"github.com/Santiago-Balcero/go-spotify/models"
	"github.com/Santiago-Balcero/go-spotify/services"
	"github.com/zmb3/spotify"
)

var client spotify.Client

func main() {
	config.LoadConfig()
	client = config.GetClient()

	http.HandleFunc("/spotify", artistStats)
	http.ListenAndServe(":8080", nil)
}

func artistStats(w http.ResponseWriter, r *http.Request) {
	response := models.CustomResponse{
		Code:    400,
		Message: "Bad request",
		Data:    nil,
	}
	artistName := r.URL.Query().Get("name")
	if artistName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		log.Println("No artist name in request - Response:", response)
		return
	}
	log.Println("Request for artist:", artistName)
	artistName = strings.ReplaceAll(artistName, "-", " ")
	artist, err := services.SearchArtist(&client, artistName)
	if err != nil {
		log.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		log.Println("Error response:", response)
		return
	}
	log.Println("Artist found:", artist)

	err = services.AnalyseArtist(&client, &artist)
	if err != nil {
		log.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		log.Println("Error response:", response)
		return
	}
	response.Code = 200
	response.Message = "Artist processed"
	response.Data = artist
	json.NewEncoder(w).Encode(response)
	log.Println("Response:", response)
}
