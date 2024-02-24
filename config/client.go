package config

import (
	"context"

	"github.com/Santiago-Balcero/go-spotify/utils"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

func GetClient() spotify.Client {
	authConfig := &clientcredentials.Config{
		ClientID:     SpotifyId,
		ClientSecret: SpotifySecret,
		TokenURL:     spotify.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())
	utils.CheckError(err)

	return spotify.Authenticator{}.NewClient(accessToken)
}
