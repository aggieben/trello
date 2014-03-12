package main

import (
	"flag"
	"fmt"
	"github.com/aggieben/trello"
	"log"
)

var userToken = flag.String("token", "", "user token obtained from trello")

func main() {
	flag.Parse()

	if *userToken == "" {
		log.Fatalf("userToken required")
	}

	client := trello.NewTrello(trello.TrelloParams{
		Version:   "1",
		AppKey:    "e9002cf6caab2a10867d22384966ab42",
		UserToken: *userToken})

	resp := <-client.Members.Me(client.Context, &trello.ModelParams{Fields: client.Members.MinimalFields()})
	fmt.Printf("got response: %v", resp)
}
