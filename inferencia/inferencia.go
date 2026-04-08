// Inferência de tipo é a capacidade do compilador de deduzir o tipo de uma variável com base no valor atribuído a ela. Em Go, isso é feito usando a sintaxe de declaração curta (:=), que permite que o compilador determine o tipo da variável automaticamente.
package main

import "fmt"

func main() {
	var nome = "Gerson"
	var idade = 44
	var versao = 3.2
	fmt.Println("O nome é", nome, "e a idade é", idade, "anos. A versão do Go é", versao)
}