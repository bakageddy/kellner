package pokemon

import (
	"github.com/carlmjohnson/requests"

	"context"
	"strings"
)

type Pokemon struct {
	Name string   `json:"name"`
	Type []string `json:"types"`
	Id   float64   `json:"id"`
}

type JsonElement = map[string]interface{}

const apiBase string = "https://pokeapi.co/api/v2/"

// Returns a new Pokemon object from the
// Pokeapi...
func New(name string) (Pokemon, error) {
	var contentBody map[string]interface{}
	var types []string
	ctx := context.Background()

	err := requests.
		URL(apiBase).
		Pathf("pokemon/%+v", name).
		ToJSON(&contentBody).
		Fetch(ctx)

	if err != nil {
		return Pokemon{}, err
	}

	id := contentBody["id"].(float64)
	mapTypes := contentBody["types"].([]interface{})
	for _, elem := range mapTypes {
		x := elem.(JsonElement)
		pokeInfo := x["type"].(JsonElement)
		pokeType := pokeInfo["name"].(string)
		placeHolder := strings.Title(strings.ToLower(pokeType))
		types = append(types, placeHolder)
	}

	return Pokemon{
		Name: name,
		Id:   id,
		Type: types,
	}, nil
}
