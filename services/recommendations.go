package services

import (
	"fmt"

	"github.com/Santiago-Balcero/go-spotify/constants"
	"github.com/Santiago-Balcero/go-spotify/utils"
	"github.com/zmb3/spotify"
)

func GetRecommendations(client *spotify.Client) {
	seed := spotify.Seeds{
		Artists: []spotify.ID{
			spotify.ID(constants.ORappaId),
			spotify.ID(constants.JoutroMundoId),
		},
		Tracks: []spotify.ID{
			spotify.ID(""), // Nao force
			spotify.ID(""), // Panthera
		},
		Genres: []string{"Funk"},
	}

	trackAttributes := spotify.NewTrackAttributes()
	trackAttributes.MinTempo(120)
	trackAttributes.MaxTempo(150)
	trackAttributes.MinDanceability(0.7)
	trackAttributes.MinEnergy(0.6)

	country := "BR"
	options := spotify.Options{
		Country: &country,
	}

	recommendations, err := client.GetRecommendations(
		seed,
		trackAttributes,
		&options,
	)
	utils.CheckError(err)

	for _, track := range recommendations.Tracks {
		fmt.Printf(
			"Recommendation: %s - %v\n",
			track.Name,
			track.ExternalURLs["spotify"],
		)
	}
}
