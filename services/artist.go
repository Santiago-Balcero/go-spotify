package services

import (
	"errors"
	"fmt"

	"github.com/Santiago-Balcero/go-spotify/models"
	"github.com/Santiago-Balcero/go-spotify/utils"
	"github.com/zmb3/spotify"
)

func AnalyseArtist(client *spotify.Client, artistId string) {
	var maxDanceability float32
	var maxDanceabilityTrack string
	var maxEnergy float32
	var maxEnergyTrack string

	artist, err := client.GetArtist((spotify.ID(artistId)))
	utils.CheckError(err)

	fmt.Println("Artist name:", artist.Name)
	fmt.Println("Artist popularity:", artist.Popularity)

	albums, err := client.GetArtistAlbums(spotify.ID(artistId))
	utils.CheckError(err)

	for _, album := range albums.Albums {
		fmt.Println("Artist album:", album.Name)
		tracks, err := client.GetAlbumTracks(album.ID)
		utils.CheckError(err)

		for _, track := range tracks.Tracks {
			if utils.ArtistInList(track.Artists, artist.Name) {
				fmt.Println("\tTrack:", track.Name)
				danceability, energy := AnalyseTrack(client, track.ID)
				if danceability > maxDanceability {
					maxDanceability = danceability
					maxDanceabilityTrack = track.Name
				}
				if energy > maxEnergy {
					maxEnergy = energy
					maxEnergyTrack = track.Name
				}
			}
		}
	}

	fmt.Println("Most danceable track:", maxDanceabilityTrack, maxDanceability)
	fmt.Println("Most energetic track:", maxEnergyTrack, maxEnergy)
}

func SearchArtist(client *spotify.Client, artistName string) (models.Artist, error) {
	result, err := client.Search(artistName, spotify.SearchTypeArtist)
	utils.CheckError(err)

	artist := models.Artist{}
	artistFound := false

	for _, a := range result.Artists.Artists {
		name := utils.ClearString(a.Name)
		if name == artistName {
			artist.Id = a.ID.String()
			artist.Name = a.Name
			artist.Popularity = a.Popularity
			artist.Genres = a.Genres
			artist.Url = a.ExternalURLs["spotify"]
			artist.Followers = int(a.Followers.Count)
			artist.Image = string(a.Images[2].URL)
			fmt.Println("Artist found: ", artist.Name, artist.Url)
			artistFound = true
			break
		}
	}

	if !artistFound {
		return artist, errors.New("artist not found")
	}

	return artist, nil
}
