package services

import (
	"fmt"
	"strings"

	"github.com/Santiago-Balcero/go-spotify/models"
	"github.com/Santiago-Balcero/go-spotify/utils"
	"github.com/zmb3/spotify"
)

func AnalyseArtist(client *spotify.Client, artistData *models.Artist) error {
	artist, err := client.GetArtist((spotify.ID(artistData.Id)))
	if err != nil {
		return fmt.Errorf("error in GetArtist: %v", err)
	}

	albums, err := client.GetArtistAlbums(spotify.ID(artistData.Id))
	if err != nil {
		return fmt.Errorf("error in GetArtistAlbums: %v", err)
	}

	albumsData := []models.Album{}

	for _, album := range albums.Albums {
		artistData.AlbumsCount++
		fmt.Println("Artist album:", album.Name)
		tracks, err := client.GetAlbumTracks(album.ID)
		utils.CheckError(err)
		albumData := models.Album{
			Name:        album.Name,
			Type:        album.AlbumType,
			ReleaseDate: album.ReleaseDate,
			Tracks:      []models.Track{},
		}
		for _, track := range tracks.Tracks {
			if utils.ArtistInList(track.Artists, artist.Name) {
				artistData.TracksCount++
				fmt.Println("\tTrack:", track.Name)
				trackData := models.Track{
					Id:   track.ID.String(),
					Name: track.Name,
				}
				err := AnalyseTrack(client, &trackData)
				if err != nil {
					return fmt.Errorf("error in AnalyseTrack: %v", err)
				}
				checkAlbumMaximums(&trackData, &albumData)
				albumData.Tracks = append(albumData.Tracks, trackData)
			}
		}
		albumsData = append(albumsData, albumData)
		checkArtistMaximums(&albumData, artistData)
	}

	artistData.Albums = albumsData
	return nil
}

func SearchArtist(client *spotify.Client, artistName string) (models.Artist, error) {
	result, err := client.Search(artistName, spotify.SearchTypeArtist)
	utils.CheckError(err)

	artist := models.Artist{}
	artistFound := false

	for _, a := range result.Artists.Artists {
		name := utils.ClearString(a.Name)
		if name == strings.ReplaceAll(artistName, " ", "") {
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
		return artist, fmt.Errorf("artist not found")
	}

	return artist, nil
}

func checkAlbumMaximums(trackData *models.Track, albumData *models.Album) {
	if trackData.Danceability > albumData.MaxDanceability {
		albumData.MaxDanceability = trackData.Danceability
		albumData.MaxDanceabilityTrack = trackData.Name
	}
	if trackData.Energy > albumData.MaxEnergy {
		albumData.MaxEnergy = trackData.Energy
		albumData.MaxEnergyTrack = trackData.Name
	}
}

func checkArtistMaximums(albumData *models.Album, artistData *models.Artist) {
	if albumData.MaxDanceability > artistData.MaxDanceability {
		artistData.MaxDanceability = albumData.MaxDanceability
		artistData.MaxDanceabilityTrack = albumData.MaxDanceabilityTrack
	}
	if albumData.MaxEnergy > artistData.MaxEnergy {
		artistData.MaxEnergy = albumData.MaxEnergy
		artistData.MaxEnergyTrack = albumData.MaxEnergyTrack
	}
}
