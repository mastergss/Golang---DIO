// Continue é uma instrução de controle de fluxo que permite pular a iteração atual de um loop e continuar com a próxima iteração. No exemplo abaixo, o loop for percorre os números de 0 a 19, mas quando encontra um número par, ele usa a instrução continue para pular a impressão desse número e continuar com o próximo número. Assim, apenas os números ímpares serão impressos.
package main

import "fmt"

func main() {
	// Loop com continue para pular números pares
	for x:= 0; x < 20; x++ {
		// Verificar se x é par
		if x % 2 == 0 {
			continue
		}
		fmt.Println("valor de x: ", x)
	}
}