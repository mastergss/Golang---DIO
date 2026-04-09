// O que é conversão de tipos em Go?
// A conversão de tipos é o processo de transformar um valor de um tipo de dado para outro tipo de dado.
// Em Go, a conversão de tipos é feita explicitamente.
// Para converter um int para float64, usamos a sintaxe: float64(a) - 'a' é a variável que queremos converter.
package main

import "fmt"
// Declarando as variáveis 'a' e 'b' escopo global
var a int = 20
var b string = "nome"

func main() {
	fmt.Println("Valor de a:", a)
	fmt.Println("Valor de b:", b)

	// Convertendo 'var a' para float64
	var c float64 = float64(a)
	fmt.Println("Valor de c (float64):", c)

	fmt.Printf("O valor de c é: %g e o valor de b é: %s", c, b)
}