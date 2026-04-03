//Exibir números entre 1 e 100 que sejam múltiplos de 3
package main

import "fmt"

func main() {
	for num := 1; num <= 100; num++ {
		if num%3 == 0 {
			fmt.Print(num, " ")
		}
	}
}