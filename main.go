package main

import (
	"log"
	"os"

	"github.com/Mercwri/crystarium/client"
)

func main() {
	_, err := client.NewCrystarium(client.CrystamiumConfig{
		ClientSecret: os.Getenv("FFLOGS_SECRET"),
		ClientID:     os.Getenv("FFLOGS_CLIENT"),
	})
	if err != nil {
		log.Panic(err)
	}
}
