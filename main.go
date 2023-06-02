// kellner
package main

import (
	"encoding/json"
	// "io/ioutil"
	"log"
	"net/http"

	"github.com/gdnand/kellner/pokemon"
)

// PokemonHandle for handling /pokemon endpoint
type PokemonHandle struct{}

// Handle responses to serve pokemon endpoint!
func (*PokemonHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	enc := json.NewEncoder(w)

	name := r.FormValue("name")
	pokemon, err := pokemon.New(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to find pokemon: ", r.RequestURI)
		return
	}

	enc.Encode(pokemon)
}

func main() {
	http.Handle("/pokemon", &PokemonHandle{})
	http.ListenAndServe(":8080", nil)
}
