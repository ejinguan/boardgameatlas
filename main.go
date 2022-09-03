package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/ejinguan/boardgameatlas/api"
)

// go run .
// go build .
// go build -o bga .
func main() {
	// bga --query "ticket to ride" --clientId abc123 --skip 10 --limit 5
	// Define the command line arguments
	query := flag.String("query", "", "Boardgame name to search")
	clientId := flag.String("clientId", "", "My BGA Client ID")
	skip := flag.Uint("skip", 0, "Skips the number of results provided.")
	limit := flag.Uint("limit", 10, "Limits the number of results returned.")

	// Parse the command line arguments
	flag.Parse()

	fmt.Println("hello world")

	// Check command line arguments
	if isNull(*query) {
		log.Fatalln("Please use --query to set the boardgame name to search")
	}
	if isNull((*clientId)) {
		log.Fatalln("Please use --clientId to set your BGA clientId")
	}

	fmt.Printf("query=%s, clientId=%s, skip=%d, limit=%d\n", *query, *clientId, *skip, *limit)

	// Instantiate a BoardgameAtlas struct client
	bga := api.New(*clientId)

	// Make the invocation
	result, err := bga.Search(context.Background(), *query, *limit, *skip)
	if nil != err {
		log.Fatalf("Cannot search for boardgame: %v", err)
	}

	// Looping through result games
	for _, g := range result.Games {
		fmt.Printf("Name: %s\n", g.Name)
		fmt.Printf("Description: %s\n", g.Description)
		fmt.Printf("URL: %s\n\n", g.Url)
	}
}

func isNull(s string) bool {
	return len(strings.TrimSpace(s)) <= 0
}
