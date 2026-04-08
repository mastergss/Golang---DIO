package main

import "fmt"

/* Converte a temperatura de Fahrenheit para Celsius 
ponto de ebulição da água */
const temperaturaFahrenheit = 212.0

func main() {
	tempF := temperaturaFahrenheit
	tempC := (tempF - 32) * 5 / 9
	// Imprime os resultados formatados e mostra os tipos de dados
	fmt.Printf("A temperatura de ebulição da água em °F é: %g(%T)\nA temperatura de ebulição da água em °C é: %g(%T)", tempF, tempF, tempC, tempC)
}