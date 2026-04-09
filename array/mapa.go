//o que é um mapa? um mapa é uma coleção de pares chave-valor, onde cada chave é única e está associada a um valor. os mapas são usados para armazenar dados de forma eficiente e permitem acesso rápido aos valores com base nas chaves. em go, os mapas são implementados como tabelas hash, o que significa que eles oferecem desempenho constante para operações de inserção, exclusão e busca.
package main

import "fmt"

func main() {
	//criando um mapa de string para int
	idades := make(map[string]int)
	idades["Alice"] = 30
	idades["Bob"] = 25
	idades["Charlie"] = 35

	fmt.Println(idades)
	//acessando um valor do mapa
	fmt.Println("Idade de Alice:", idades["Alice"])
}