package main

import "fmt"

/* Converte a temperatura de Fahrenheit para Celsius 
ponto de ebulição da água */
const temperaturaFahrenheit = 212.0

func main() {
	var tempF = temperaturaFahrenheit
	var tempC = (tempF - 32) * 5 / 9
	fmt.Println("A temperatura de ebulição da água em °F é: ", tempF)
	fmt.Println("A temperatura de ebulição da água em °C é: ", tempC)
}