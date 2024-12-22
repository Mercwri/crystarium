package main

import (
	"log"
	"os"

	"github.com/Mercwri/crystarium/client"
)

func main() {
	c, err := client.NewCrystarium(client.CrystamiumConfig{
		ClientSecret: os.Getenv("FFLOGS_SECRET"),
		ClientID:     os.Getenv("FFLOGS_CLIENT"),
	})
	if err != nil {
		log.Panic(err)
	}
	user, err := c.GetUser(503278)
	if err != nil {
		log.Panic(err)
	}
	log.Println(user.UserData.User.Name)
}
