package client

import (
	"context"

	userdata "github.com/Mercwri/crystarium/queries/UserData"
	"github.com/hasura/go-graphql-client"
	"golang.org/x/oauth2/clientcredentials"
)

type Cystarium struct {
	Config CrystamiumConfig
	QGL    *graphql.Client
}

type CrystamiumConfig struct {
	ClientSecret string
	ClientID     string
}

func NewCrystarium(config CrystamiumConfig) (Cystarium, error) {
	clientConfig := clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		TokenURL:     "https://www.fflogs.com/oauth/token",
	}
	client := clientConfig.Client(context.Background())
	gql := graphql.NewClient("https://www.fflogs.com/api/v2/client", client)
	return Cystarium{
		Config: config,
		QGL:    gql,
	}, nil
}

func (c *Cystarium) GetUser(id int) (userdata.Query, error) {
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
