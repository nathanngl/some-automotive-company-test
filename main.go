package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	var pokemons []map[string]interface{}

	for i := 1; i <= 20; i++ {
		poke := getPokemon(i)

		pokemons = append(pokemons, poke)
	}

	for i := 0; i < len(pokemons)-1; i++ {
		if pokemons[i]["attackStats"].(float64) > pokemons[i+1]["attackStats"].(float64) {
			pokemons[i], pokemons[i+1] = pokemons[i+1], pokemons[i]
		}
	}

	for i := 0; i < len(pokemons); i++ {
		fmt.Println(pokemons[i]["attackStats"], pokemons[i]["speciesName"])
	}
}

func getPokemon(id int) map[string]interface{} {
	url := "https://pokeapi.co/api/v2/pokemon"

	client := http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/%d", url, id), nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	responseBOdy, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	responseReturn := string(responseBOdy)

	var data map[string]interface{}
	err = json.Unmarshal([]byte(responseReturn), &data)
	if err != nil {
		panic(err)
	}

	species := data["species"].(map[string]interface{})
	speciesName := species["name"].(string)

	stats := data["stats"].([]interface{})
	attackStats := float64(0)
	for i, v := range stats {
		if i == 0 {
			base := v.(map[string]interface{})
			attackStats = base["base_stat"].(float64)
			break
		}
	}

	result := map[string]interface{}{
		"speciesName": speciesName,
		"attackStats": attackStats,
	}

	return result
}
