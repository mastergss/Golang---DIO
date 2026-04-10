//Recursão é uma técnica de programação onde uma função chama a si mesma para resolver um problema. Em Go, as funções podem ser recursivas, ou seja, podem chamar a si mesmas para resolver um problema de forma mais simples e elegante. A recursão é frequentemente usada para resolver problemas que podem ser divididos em subproblemas menores, como a série de Fibonacci, a fatoração de números, entre outros.

package main

import "fmt"

func fatorial(x int) int {
	if x == 0 { //caso base: o fatorial de 0 é 1
		return 1
	}
	fmt.Println("Calculando fatorial de ", x) //imprime o valor de x para mostrar a recursão
	return x * fatorial(x-1) //caso recursivo: x! = x * (x-1)!
}

func main() {
	fmt.Println("Fatorial de 5: ", fatorial(5)) //chama a função fatorial e imprime o resultado
}