package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ejinguan/boardgameatlas/api"
	"github.com/fatih/color"
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
	timeout := flag.Uint("timeout", 10, "Timeout")

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout*uint(time.Second)))
	defer cancel() // Defer to end of program, will cancel any pending request when timeout happens or main() ends

	result, err := bga.Search(ctx, *query, *limit, *skip)
	if nil != err {
		log.Fatalf("Cannot search for boardgame: %v", err)
	}

	// Colors
	boldgreen := color.New(color.Bold).Add(color.FgHiGreen).SprintFunc()

	// Looping through result games
	for _, g := range result.Games { // Ignoring the iterator variable, just get the game
		fmt.Printf("%s %s\n", boldgreen("Name"), g.Name)
		fmt.Printf("%s: %s\n", boldgreen("Description"), g.Description)
		fmt.Printf("%s: %s\n\n", boldgreen("URL"), g.Url)
	}
}

func isNull(s string) bool {
	return len(strings.TrimSpace(s)) <= 0
}
