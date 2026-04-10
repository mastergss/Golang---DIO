//Ponteiro é uma variável que armazena o endereço de memória de outra variável. Em Go, o operador & é usado para obter o endereço de memória de uma variável, e o operador * é usado para acessar o valor armazenado no endereço de memória.
//Como utilizar ponteiros em Go? Para usar ponteiros em Go, você pode declarar uma variável do tipo ponteiro usando o operador *, e depois usar o operador & para obter o endereço de memória de uma variável e atribuí-lo ao ponteiro. Você também pode usar o operador * para acessar o valor armazenado no endereço de memória apontado pelo ponteiro. Por exemplo:

package main

import "fmt"

func inicial (x int ,yPtr *int) {
	x = 0 //a variável x é uma cópia da variável original, então a alteração não afeta a variável original
	fmt.Println("Valor de x dentro da função: ", x) //imprime o valor de x dentro da função
	*yPtr = 5
	fmt.Println("Valor de y dentro da função: ", *yPtr) //imprime o valor de y dentro da função, que foi alterado usando o ponteiro
}

func main() {
	x := 10
	y := 15
	fmt.Println("Valor de x antes da função: ", x) //imprime o valor de x antes de chamar a função
	fmt.Println("Valor de y antes da função: ", y)
	inicial(x, &y)
	fmt.Println("Valor de x depois da função: ", x) //imprime o valor de x depois de chamar a função, que não foi alterado
	fmt.Println("Valor de y depois da função: ", y) //imprime o valor de y depois de chamar a função, que não foi alterado
}