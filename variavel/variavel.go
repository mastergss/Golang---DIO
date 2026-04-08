//Declaração de variável em Go
//Em Go, as variáveis são declaradas usando a palavra-chave "var" seguida pelo nome da variável, o tipo e o valor opcional.
package main

import "fmt"

func main() {
	// Declaração de variável com tipo explícito
	var x string = "Hello, World!"
	fmt.Println(x)
	// Declaração de variável com inferência de tipo
	var y = 42
	fmt.Println(y)
	// Declaração de variável com inferência de tipo e valor opcional
	var z = "Go"
	fmt.Println(z)
	// Declaração de múltiplas variáveis
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)
}