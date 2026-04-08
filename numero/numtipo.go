//Exemplo de declaração de variáveis numéricas em Go
//Neste exemplo, declaramos duas variáveis numéricas: x do tipo int e y do tipo float64. Em seguida, imprimimos os valores dessas variáveis usando fmt.Println e fmt.Printf para mostrar o tipo de cada variável.
package main

import "fmt"

var x int = 10
var y float64 = 20.5

func main () {
	fmt.Println("Valor de x: ", x) //numero inteiro int
	fmt.Println("Valor de y: ", y) //numero decimal float64
	fmt.Printf("x = %d,%T e y = %g,%T", x, x, y, y) //imprime os valores de x e y usando formatação
}