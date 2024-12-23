package client

import (
	"context"
	"log"

	characterdata "github.com/Mercwri/crystarium/queries/CharacterData"
	reportdata "github.com/Mercwri/crystarium/queries/ReportData"
	userdata "github.com/Mercwri/crystarium/queries/UserData"
	"github.com/hasura/go-graphql-client"
	"golang.org/x/oauth2/clientcredentials"
)

type Crystarium struct {
	Config CrystamiumConfig
	QGL    *graphql.Client
}

type CrystamiumConfig struct {
	ClientSecret string
	ClientID     string
}

func NewCrystarium(config CrystamiumConfig) (Crystarium, error) {
	clientConfig := clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		TokenURL:     "https://www.fflogs.com/oauth/token",
	}
	client := clientConfig.Client(context.Background())
	gql := graphql.NewClient("https://www.fflogs.com/api/v2/client", client)
	// gql = gql.WithDebug(true)
	return Crystarium{
		Config: config,
		QGL:    gql,
	}, nil
}

func (c *Crystarium) GetUser(id int) (userdata.Query, error) {
	var query userdata.Query
	vars := map[string]interface{}{
		"id": graphql.Int(id),
	}
	err := c.QGL.Query(context.Background(), &query, vars)
	if err != nil {
		return query, err
	}
	return query, nil
}

func (c *Crystarium) GetCharacter(name string, server string, region string) (characterdata.Character, error) {
	var query characterdata.Query
	vars := map[string]interface{}{
		"name":   graphql.String(name),
		"server": graphql.String(server),
		"region": graphql.String(region),
		"zoneID": graphql.Int(58),
	}
	err := c.QGL.Query(context.Background(), &query, vars)
	if err != nil {
		return query.CharacterData.Character, err
	}
	return query.CharacterData.Character, nil
}

func (c *Crystarium) GetReport(code string) (reportdata.Report, error) {
	var query reportdata.ReportQuery
	vars := map[string]interface{}{
		"code": graphql.String(code),
	}
	err := c.QGL.Query(context.Background(), &query, vars)
	if err != nil {
		return query.ReportData.Report, err
	}
	return query.ReportData.Report, nil
}

func (c *Crystarium) GetFightEvents(report reportdata.Report, id int) ([]reportdata.Data, error) {
	var events []reportdata.Data
	var query reportdata.FightQuery
	vars := map[string]interface{}{
		"code":      graphql.String(report.Code),
		"fightID":   []graphql.Int{graphql.Int(id)},
		"fex":       graphql.String("type!=\"combatantinfo\""),
		"timestamp": graphql.Float(0),
	}
	err := c.QGL.Query(context.Background(), &query, vars)
	if err != nil {
		return query.ReportData.Report.Events.Data, err
	}
	nxp := query.ReportData.Report.Events.NextPageTimestamp
	events = append(events, query.ReportData.Report.Events.Data...)
	for nxp != 0 {
		var queryP reportdata.FightQuery
		log.Println(nxp)
		vars["timestamp"] = nxp
		err := c.QGL.Query(context.Background(), &queryP, vars)
		if err != nil {
			return events, err
		}
		events = append(events, queryP.ReportData.Report.Events.Data...)
		nxp = queryP.ReportData.Report.Events.NextPageTimestamp
	}
	return events, err
}
