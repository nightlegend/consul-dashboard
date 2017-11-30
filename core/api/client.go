package api

import (
	"github.com/hashicorp/consul/api"
)

// NewCli is new a consul client.
func NewCli() *api.Client {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	return client
}
