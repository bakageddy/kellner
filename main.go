package main

import (
	"context"
	"log"
	"strings"

	"github.com/carlmjohnson/requests"
)

type JsonElement = map[string]interface{}

const apiBase string = "https://pokeapi.co/api/v2/"

func main() {
	var contentBody map[string]interface{}
	ctx := context.Background()

	err := requests.
		URL(apiBase).
		Path("pokemon/charizard").
		ToJSON(&contentBody).
		Fetch(ctx)

	
	if err != nil {
		log.Fatalln(err)
	}

	types := contentBody["types"].([]interface{})
	for _, elem := range types {
		x := elem.(JsonElement)
		pokeInfo := x["type"].(JsonElement)
		pokeType := pokeInfo["name"].(string)
		log.Println(strings.Title(strings.ToLower(pokeType)))
	}
}
