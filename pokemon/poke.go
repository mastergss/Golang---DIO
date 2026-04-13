//Usaremos a pokeapi - tudo sobre pokemon
//Para consultar: http://pokeapi.com/api/v2/pokedex/kanto/
//Verbos em HTTP:
// GET: lista registros
// POST: adiciona um novo registro
// DELETE: remove um registro
// PUT e PATCH: editar um registro

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	//response utilizado para mapear 'site'
	response, err := http.Get("http://pokeapi.com/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	//responseData vai armazenar conteúdo do body (em bytes) vindo do response
	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	//utilizar string para converter bytes em string
	fmt.Println(string(responseData))
}