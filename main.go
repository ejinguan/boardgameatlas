package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

// go run .
// go build .
// go build -o bga .
func main() {
	// bga --query "ticket to ride" --clientID abc123 --skip 10 --limit 5
	// Define the command line arguments
	query := flag.String("query", "", "Boardgame name to search")
	clientID := flag.String("clientID", "", "My BGA Client ID")
	skip := flag.Uint("skip", 0, "Skips the number of results provided.")
	limit := flag.Uint("limit", 10, "Limits the number of results returned.")

	// Parse the command line arguments
	flag.Parse()

	fmt.Println("hello world")

	// Check command line arguments
	if isNull(*query) {
		log.Fatalln("Please use --query to set the boardgame name to search")
	}
	if isNull((*clientID)) {
		log.Fatalln("Please use --clientID to set your BGA clientID")
	}

	fmt.Printf("query=%s, clientId=%s, skip=%d, limit=%d\n", *query, *clientID, *skip, *limit)
}

func isNull(s string) bool {
	return len(strings.TrimSpace(s)) <= 0
}
