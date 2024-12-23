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
	report, err := c.GetReport("k7vRT9ZpgzNVCXtQ")
	if err != nil {
		log.Panic(err)
	}
	log.Println(report)
	e, err := c.GetFightEvents(report, 3)
	if err != nil {
		log.Panic(err)
	}
	sum := 0
	for i, ev := range e {
		log.Println(i, ev.Timestamp, ev.Type, ev.SourceID, ev.TargetID, ev.Amount, ev.FinalizedAmount)
		if ev.Type == "damage" && ev.TargetID == 11 {
			sum = sum + int(ev.Amount)
		}
	}
	log.Print(sum)
}
