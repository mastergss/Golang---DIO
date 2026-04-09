// Método é uma função associada a um tipo de dado. Ele permite que o tipo de dado execute operações específicas.
package main

import "fmt"
// Definindo um tipo de dado para representar um retângulo
type retangulo struct {
	altura, comprimento float64
}

// Método para calcular a área do retângulo
func (r retangulo) area() float64 {
	return r.altura * r.comprimento
}
// Método para calcular o perímetro do retângulo
func (r retangulo) perimetro() float64 {
	return 2 * r.altura + 2 * r.comprimento 
}

func main() {
	//declarar e inicializar um retângulo
	r := retangulo{altura: 5.5, comprimento: 10.1}
	//chamar os métodos para calcular a área e o perímetro do retângulo
	fmt.Println("Área: ", r.area())
	fmt.Println("Perimetro: ", r.perimetro())
}