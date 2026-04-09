//Brincadeira Pin & Pan
//Nos casos que o número é múltiplo de 3, exibir "Pin" ao invés do número
//Nos casos que o número é múltiplo de 5, exibir "Pan" ao invés do número
//Nos casos que o número é múltiplo de 3 e 5, exibir "PinPan" ao invés do número

package main

import "fmt"

func main() {
	// Percorre os números de 1 a 100
	for num := 1; num <= 100; num++ {
		// Verifica se o número é múltiplo de 3 e 5
		if num % 3 == 0 && num % 5 == 0 {
			// Exibe "PinPan" no console
			fmt.Println("PinPan")
			//
		} else if num % 3 == 0 {
			// Exibe "Pin" no console
			fmt.Println("Pin")
			// Verifica se o número é múltiplo de 5
		} else if num % 5 == 0 {
			// Exibe "Pan" no console
			fmt.Println("Pan")
		}
	}
}