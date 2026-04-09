// Loops hierárquicos são loops aninhados, onde um loop é colocado dentro de outro. O loop externo controla a iteração principal, enquanto o loop interno é executado para cada iteração do loop externo. Isso é útil para trabalhar com estruturas de dados multidimensionais ou para realizar tarefas que exigem múltiplos níveis de repetição.
package main

import "fmt"

func main() {
	//Loop Hierárquico
	// Loop externo para as horas
	for hora := 0; hora < 12; hora++ {
		fmt.Printf("Hora: %02d\n", hora)
		// Loop interno para os minutos
		for minuto := 0; minuto < 60; minuto++ {
			fmt.Printf("%02d:%02d\n", hora, minuto)
		}
	}
}