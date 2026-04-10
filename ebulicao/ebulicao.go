// Programa para calcular a temperatura de ebulição e fusão da água em Kelvin e Celsius
package main

import "fmt"

// Definindo as constantes para a temperatura de ebulição e fusão da água em Kelvin
const temperaturaEbuliçãoKelvin = 373.15
const temperaturaFusaoKelvin = 273.15

func main() {
	// Calculando a temperatura de ebulição da água em Kelvin para Celsius
	tempK := temperaturaEbuliçãoKelvin
	tempC := tempK - 273.15
	// Imprimindo os resultados
	fmt.Printf("A temperatura de ebulição da água em K é: %g(%T)\nA temperatura de ebulição da água em °C é: %g(%T)", tempK, tempK, tempC, tempC)

	// Calculando a temperatura de fusão da água em Kelvin para Celsius
	fusaoK := temperaturaFusaoKelvin
	fusaoC := fusaoK - 273.15
	// Imprimindo os resultados - Printf para formatar a saída, %g para o valor numérico e %T para o tipo da variável
	fmt.Printf("\nA temperatura de fusão da água em K é: %g(%T)\nA temperatura de fusão da água em °C é: %g(%T)", fusaoK, fusaoK, fusaoC, fusaoC)
}