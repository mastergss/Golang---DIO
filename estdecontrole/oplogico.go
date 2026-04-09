// Operadores lógicos são usados para combinar expressões booleanas. Os principais operadores lógicos são:
// - && (AND): Retorna true se ambas as expressões forem verdadeiras.
// - || (OR): Retorna true se pelo menos uma das expressões for verdadeira.
// - ! (NOT): Inverte o valor lógico de uma expressão.
package main

import "fmt"

func main() {
	x := 2
	// Usando o operador OR (||) para verificar se x é igual a 2 ou 3
	if (x == 2 || x == 3) {
		fmt.Println("x é igual a 2 ou 3")
	} else {
		fmt.Println("x não é igual a 2 ou 3")
	}

	y := 5
	// Usando o operador AND (&&) para verificar se y é divisível por 2 e 3
	if (y % 2 == 0 && y % 3 == 0) {
		fmt.Println("y é divisível por 2 e 3")
	} else {
		fmt.Println("y não é divisível por 2 e 3")
	}

	z := 10
	// Usando o operador NOT (!) para verificar se z não é igual a 10
	if !(z == 10) {
		fmt.Println("z não é igual a 10")
	} else {
		fmt.Println("z é igual a 10")
	}
}