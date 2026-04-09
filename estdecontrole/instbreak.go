// Break é usado para interromper um loop ou switch. Ele pode ser usado para sair de um loop infinito ou para sair de um switch case.
package main

import "fmt"

func main() {
	// Loop infinito com break
	x := 0
	for {
		if x < 20 {
			fmt.Println("valor de x: ", x)
			x++
		} else {
			// Interrompe o loop quando x for igual a 20
			break
		}
	}
}