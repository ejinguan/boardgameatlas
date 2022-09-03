package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

const SEARCH_URL = "https://api.boardgameatlas.com/api/search"

// Like a class
type BoardgameAtlas struct {
	// "members" - lowercase=private, Uppercase=public
	clientId string
}

// "Method" in BoardgameAtlas
// Pass a BGA to the Search function (receiver), Can also pass a *BGA
func (b BoardgameAtlas) Search(ctx context.Context, query string, limit uint, skip uint) error {

	// Create HTTP Client
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, SEARCH_URL, nil)
	// Check if there is any error
	if nil != err {
		// return an error object
		return fmt.Errorf("cannot create HTTP client: %v", err)
	}

	// Get the query string object
	qs := req.URL.Query()
	// Populate the URL with query params
	qs.Add("name", query)
	qs.Add("limit", fmt.Sprintf("%d", limit)) // sprintf
	qs.Add("skip", strconv.Itoa(int(skip)))   // Cast skip to integer
	qs.Add("client_id", b.clientId)
	// Encode the query params, add it back to the request
	req.URL.RawQuery = qs.Encode()

	fmt.Printf("URL = %s\n", req.URL.String())

	return nil
}

// New functions as a constructor
func New(clientId string) BoardgameAtlas {
	return BoardgameAtlas{clientId: clientId}
	// or can use
	// return BoardgameAtlas{clientId}
}
