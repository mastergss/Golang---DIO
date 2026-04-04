//Brincadeira Pin & Pan
//Nos casos que o número é múltiplo de 3, exibir "Pin" ao invés do número
//Nos casos que o número é múltiplo de 5, exibir "Pan" ao invés do número
//Nos casos que o número é múltiplo de 3 e 5, exibir "PinPan" ao invés do número

package main

import "fmt"

func main() {
	for num := 1; num <= 100; num++ {
		if num % 3 == 0 && num % 5 == 0 {
			fmt.Println("PinPan")
		} else if num % 3 == 0 {
			fmt.Println("Pin")
		} else if num % 5 == 0 {
			fmt.Println("Pan")
		}
	}
}