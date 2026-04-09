// Interface é um tipo que define um conjunto de métodos que outros tipos podem implementar.
package main

import (
	"fmt"
	"math" // Importar pacote math para usar a constante Pi e funções matemáticas
)

// Definir interface "geometria" com métodos para calcular área e perímetro
type geometria interface {
	area() float64
	perimetro() float64
}

// Definir tipo "circulo" que implementa a interface "geometria"
type circulo struct {
	raio float64
}

// Definir tipo "quadrado" que implementa a interface "geometria"
type quadrado struct {
	lado float64
}

// Implementar método area para o tipo circulo
func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}
// Implementar método perimetro para o tipo circulo
func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// Implementar método area para o tipo quadrado
func (q quadrado) area() float64 {
	return q.lado * q.lado
}
// Implementar método perimetro para o tipo quadrado
func (q quadrado) perimetro() float64 {
	return 4 * q.lado
}

// Função para calcular a área e o perímetro de uma forma geométrica
func calcular(g geometria) {
	fmt.Println("Valor:", g)
	fmt.Println("Área:", g.area())
	fmt.Println("Perímetro:", g.perimetro())
}

func main() {
	// Criar um círculo e um quadrado
	c := circulo{raio: 4}
	q := quadrado{lado: 3}

	// Calcular área e perímetro de cada forma
	fmt.Println("Círculo:")
	calcular(c)
	fmt.Println("Quadrado:")
	calcular(q)
}