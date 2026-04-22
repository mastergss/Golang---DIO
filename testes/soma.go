package main

import "fmt"

func main() {
	x := soma(1, 2, 3)
	y := multiplica (10, 10)
	fmt.Println(x, y)
}

func soma(i ...int) int{
	total := 0
	for _, valor := range i {
		total += valor
	}
	return total
}

func multiplica(i ...int) int{
	total := 1
	for _, valor := range i {
		total *= valor
	}
	return total
}