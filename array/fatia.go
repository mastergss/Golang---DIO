//o que é fatia? é um tipo de estrutura de dados que representa uma sequência de elementos do mesmo tipo. é uma referência a um segmento de um array, permitindo acessar e manipular os elementos sem precisar copiar os dados. as fatias são dinâmicas, ou seja, podem crescer ou diminuir de tamanho conforme necessário. elas são amplamente utilizadas em go para trabalhar com coleções de dados, como listas, filas e pilhas.
package main

import "fmt"

func main() {
	fatia := make([]float64, 5)
	fmt.Println(fatia)

	arr := [5]float64{8.6, 9.5, 6.2, 7.9, 3.3}
	fatia = arr[0:5]
	fmt.Println(fatia)

	fatia = arr[1:4]
	fmt.Println(fatia)

	fatia1 := []int{1, 2, 3}
	fatia2 := append(fatia1, 4, 5)
	fmt.Println(fatia1, fatia2)

	fatia3 := make([]int, 2)
	copy(fatia3, fatia2)
	fmt.Println(fatia3)
}
//var x []float64
//fatia := make([]float64, 3) fatia de float64 com 3 elementos
//fatia1 := [low:high] fatia de x, do índice low ao high-1
//fatia2 := x[0:3] fatia de x, do índice 0 ao 2
//acrescentar elementos a uma fatia: fatia = append(fatia, elemento1, elemento2, ...)
//copiar uma fatia: copy(destino, origem) destino recebe os elementos de origem, até o tamanho de destino