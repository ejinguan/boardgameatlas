package main

import (
	"flag"
	"fmt"
)

// go run .
func main() {
	// bga --query "ticket to ride" --clientId abc123 --skip 10 --limit 5
	// Define the command line arguments
	query := flag.String("query", "", "Boardgame name to search")
	clientID := flag.String("clientId", "SuzP6IsqhN", "My BGA Client ID")
	skip := flag.Uint("skip", 0, "Skips the number of results provided.")
	limit := flag.Uint("limit", 10, "Limits the number of results returned.")

	// Parse the command line arguments
	flag.Parse()

	fmt.Println("hello world")
	fmt.Printf("query=%s, clientId=%s, skip=%d, limit=%d\n", *query, *clientID, *skip, *limit)
}
