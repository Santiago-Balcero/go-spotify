package services

import (
	"fmt"

	"github.com/Santiago-Balcero/go-spotify/utils"
	"github.com/zmb3/spotify"
)

func AnalyseTrack(
	client *spotify.Client,
	trackId spotify.ID,
) (float32, float32) {
	var danceability float32
	var energy float32
	analysis, err := client.GetAudioAnalysis(trackId)
	utils.CheckError(err)

	features, err := client.GetAudioFeatures(trackId)
	utils.CheckError(err)

	fmt.Println("\t\tLoudness:", analysis.Track.Loudness)
	fmt.Println("\t\tTempo:", analysis.Track.Tempo)

	for _, feat := range features {
		fmt.Println("\t\tAcousticness:", feat.Acousticness)
		fmt.Println("\t\tDanceability:", feat.Danceability)
		fmt.Println("\t\tEnergy:", feat.Energy)
		fmt.Println("\t\tLoudness:", feat.Loudness)
		fmt.Println("\t\tLiveness:", feat.Liveness)
		fmt.Println("\t\tInstrumentalness:", feat.Instrumentalness)
		danceability = feat.Danceability
		energy = feat.Energy
	}
	return danceability, energy
}
