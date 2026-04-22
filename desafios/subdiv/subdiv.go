// Desafio: escrever e testar funções de subtração e divisão
package main

import "fmt"

func main() {
	sub := subtrair(10, 2)
	div := dividir(4, 2)
	fmt.Println(sub, div)
}

func subtrair(i ...int) int{
	total := 20
	for _, valor := range i {
		total -= valor
	}
	return total
}

func dividir(i ...int) int{
	total := 8
	for _, valor := range i {
		total /= valor
	}
	return total
}