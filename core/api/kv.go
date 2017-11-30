package api

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

// Get a handle to the KV API
var cli = NewCli()
var kv = cli.KV()

// Put a key value to kv store.
func Put(key string, value string) bool {
	// PUT a new KV pair
	p := &api.KVPair{Key: key, Value: []byte(value)}
	_, err := kv.Put(p, nil)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Get a key from store .
func Get(key string) string {
	// Lookup the pair
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(pair.Value)
}
