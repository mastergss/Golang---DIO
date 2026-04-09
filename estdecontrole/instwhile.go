//While é uma estrutura de controle de fluxo que permite executar um bloco de código repetidamente enquanto uma condição for verdadeira. Em Go, não existe uma estrutura de controle específica para while, mas podemos usar o for para simular o comportamento do while.
package main

import "fmt"

func main() {
	//utilizar o for para imprimir os números de 1 a 5
	for i := 1; i <= 5; i++ {
		// imprimir a condição de i
		fmt.Println(i, " é menor ou igual a 5")
	}

	//utilizar o for como while para imprimir os números de 1 a 5
	y := 1
	// enquanto y for menor ou igual a 5, imprimir a condição de y
	for y <= 5 {
		fmt.Println(y, " y é menor ou igual a 5")
		y++
	}
}