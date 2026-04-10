//Pacote sort fornece funções para ordenar slices e user-defined collections. Ele inclui funções para ordenar slices de tipos básicos, como int, float64 e string, bem como funções para ordenar slices de tipos personalizados usando interfaces. O pacote sort é amplamente utilizado para organizar dados em Go, permitindo que os desenvolvedores classifiquem e ordenem seus dados de forma eficiente em seus programas.

//Exemplo de uso do pacote sort:

package main

import (
	"fmt"
	"sort"
)
// Definindo um tipo personalizado Pessoa com campos Nome e Idade
type Pessoa struct {
	Nome string
	Idade int
}

// PorIdade é um tipo que implementa a interface sort.Interface para ordenar um slice de Pessoa por idade
type PorIdade []Pessoa
// Len retorna o número de elementos na coleção PorIdade
func (pessoa PorIdade) Len() int {
	return len(pessoa)
}
// Less compara a idade de duas pessoas e retorna true se a idade da pessoa i for menor que a idade da pessoa j
func (pessoa PorIdade) Less(i, j int) bool {
	return pessoa[i].Idade < pessoa[j].Idade
}
// Swap troca as posições de duas pessoas na coleção PorIdade
func (pessoa PorIdade) Swap(i, j int) {
	pessoa[i], pessoa[j] = pessoa[j], pessoa[i]
}

func main() {
	// Criando um slice de people com diferentes idades
	people := []Pessoa{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	// Ordenando as pessoas por idade usando a função sort.Sort e a interface PorIdade
	sort.Sort(PorIdade(people))
	fmt.Println("Pessoas ordenadas por idade:", people)
}