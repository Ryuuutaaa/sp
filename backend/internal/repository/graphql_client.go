package repository

import (
	"context"

	"github.com/machinebox/graphql"
)

type GQLClient struct {
	client *graphql.Client
}

func NewGQLClient(url string) *GQLClient {
	return &GQLClient{
		client: graphql.NewClient(url),
	}
}

func (g *GQLClient) Run(query string, vars map[string]interface{}, resp interface{}) error {
	req := graphql.NewRequest(query)
	for k, v := range vars {
		req.Var(k, v)
	}
	return g.client.Run(context.Background(), req, resp)
}
