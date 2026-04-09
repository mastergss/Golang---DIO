// Arrays são estruturas de dados que armazenam uma coleção de elementos do mesmo tipo. Eles têm um tamanho fixo, o que significa que o número de elementos deve ser definido no momento da declaração. 
// Os arrays são indexados, o que permite acessar cada elemento usando um índice numérico. 
// Em Go, os arrays são valores, o que significa que quando você atribui um array a outra variável ou passa um array para uma função, uma cópia do array é criada. 
// Isso pode levar a um comportamento inesperado se você não estiver ciente disso, especialmente ao trabalhar com grandes arrays.
package main

import "fmt"

func main() {
	var arr[5] float64
	arr[0] = 8.6
	arr[1] = 9.5
	arr[2] = 6.2
	arr[3] = 7.9
	arr[4] = 3.3
	//
	total := 0.0
	//for i := 0; i < 5; i++ { substiruir 5 por len(arr) para evitar erros de tipo
	for i := 0; i < len(arr); i++ { 
		fmt.Printf("Nota %d: %.2f\n", i+1, arr[i])
		total += arr[i]
	}
	fmt.Printf("Total: %.2f\n", total)
	//fmt.Println("Média: ", total/5)
	//fmt.Println("Média: ", total/len(arr)) corrigindo o erro de tipo
	fmt.Println("Média: ", total/float64(len(arr)))
}