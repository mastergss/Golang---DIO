//Usaremos a pokeapi - tudo sobre pokemon
//Para consultar: http://pokeapi.co/api/v2/pokedex/kanto/
//Verbos em HTTP:
// GET: lista registros
// POST: adiciona um novo registro
// DELETE: remove um registro
// PUT e PATCH: editar um registro
// Criar 3 estruturas para consumir dados: response; pokemon; pokemonSpecies

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Struct principal, consolida dados da 'Specie' e 'Pokemon'
type Response struct {
	Name string 		`json:"name"`
	Pokemon []Pokemon	`json:"pokemon_entries"`
}

// Captura dados da Pokedex
type Pokemon struct {
	EntryNumber int			`json:"entry_number"`
	Specie PokemonSpecies	`json:"pokemon_species"`
}

// Captura dados da espécie
type PokemonSpecies struct {
	Name string `json:"name"`
}

func main() {
	//response utilizado para mapear 'site', fazer requisão HTTP
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// Aguardar término da função principal e fechar requisição
	defer response.Body.Close()

	// responseData vai armazenar conteúdo do body (em bytes) vindo do response
	responseData, err := io.ReadAll(response.Body)
	// Em caso de erro na leitura, imprime erro
	if err != nil {
		log.Fatal(err)
	}

	//utilizar string para converter bytes em string, imprimir JSON bruto
	fmt.Println(string(responseData))

	// responseObject armazenar dados do JSON convertidos da Struct 'Response'
	var responseObject Response
	// desempacotar objeto JSON e tratar possível erro
	err = json.Unmarshal(responseData, &responseObject)

	if err != nil {
		log.Fatal("Erro ao fazer parse do JSON: ", err)
	}

	// Exibir nome da Pokedex
	fmt.Println("Pokedex: ", responseObject.Name)

	// Listar Pokemons
	for _, pokemon := range responseObject.Pokemon {
		fmt.Println(pokemon.EntryNumber)
		fmt.Println(pokemon.Specie)
		fmt.Println(pokemon.Specie.Name)
		fmt.Printf("Número: %d Espécie: %s Nome: %s\n", pokemon.EntryNumber, pokemon.Specie, pokemon.Specie.Name)
	}
}