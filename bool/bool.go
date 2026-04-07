//Valores booleanos são usados para representar valores de verdade, ou seja, verdadeiro (true) ou falso (false). Eles são amplamente utilizados em programação para controle de fluxo, tomada de decisões e expressões condicionais.
package main

import "fmt"

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)
}