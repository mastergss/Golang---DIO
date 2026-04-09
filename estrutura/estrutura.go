// Estrutura é um tipo de dado composto que agrupa variáveis de diferentes tipos sob um mesmo nome. Ela é usada para representar objetos do mundo real, como pessoas, carros, etc.

package main

import "fmt"
// Definindo a estrutura Pessoa
type pessoa struct {
	nome string
	idade int
}

func main() {
	// Criando instâncias da estrutura Pessoa
	p1 := pessoa{nome: "João", idade: 30}
	p2 := pessoa{nome: "Maria", idade: 25}
	// Imprimindo as informações das pessoas
	fmt.Println("Pessoa 1:", p1)
	fmt.Println("Pessoa 2:", p2)
	// Acessando os campos da estrutura
	fmt.Println(pessoa{"Ana", 54})
	// Acessando o campo nome da estrutura p1
	fmt.Println(p1.nome)
}