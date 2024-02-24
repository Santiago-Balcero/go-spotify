package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Santiago-Balcero/go-spotify/config"
	"github.com/Santiago-Balcero/go-spotify/services"
	"github.com/Santiago-Balcero/go-spotify/utils"
)

func main() {
	fmt.Println("¡SPOTIFY!")

	config.LoadConfig()

	client := config.GetClient()

	var artistName string

	fmt.Println("Client ready")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter artist name: ")
	artistName, err := reader.ReadString('\n')
	utils.CheckError(err)
	artistName = utils.ClearString(artistName)

	artist, err := services.SearchArtist(&client, artistName)
	utils.CheckError(err)

	var confirm string
	fmt.Println("Check the given URL to verify artist")
	fmt.Print("Confirm artist? [y/n]: ")
	fmt.Scan(&confirm)
	confirm = utils.ClearString(confirm)

	if confirm != "y" {
		return
	}

	services.AnalyseArtist(&client, &artist)
}
