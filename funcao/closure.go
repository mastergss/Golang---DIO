//Closure é uma função que "captura" variáveis do escopo externo, mesmo após o escopo externo ter terminado.
//Em Go, as closures são criadas quando uma função é definida dentro de outra função e faz referência a variáveis do escopo externo.
package main

import "fmt"

func main() {
	x := 0 //variável do escopo externo
	//definindo a closure que incrementa a variável x e retorna o novo valor
	incrementar := func() int{
		x++ //a closure captura a variável x do escopo externo
		return x
	}
	fmt.Println(incrementar()) //chama a closure e imprime o resultado
	fmt.Println(incrementar()) //chama a closure novamente e imprime o resultado

	y := incrementar()
	fmt.Println(y) //imprime o valor de y, que é o resultado da última chamada da closure
}