package main

import "fmt"

func main() {
	//criando um mapa de string para int
	x := make(map[string]int)
	x["chave"] = 30
	//acessando um valor do mapa
	fmt.Println("Valor da chave:", x["chave"])
	//criando um mapa de int para int
	y := make(map[int]int)
	y[1] = 10
	y[2] = 20
	fmt.Println(y[1], y[2])
	//criando um mapa de string para string
	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["O"] = "Oxigênio"
	elemento["C"] = "Carbono"
	fmt.Println(elemento["O"])
	
}
//o que é um mapa? um mapa é uma coleção de pares chave-valor, onde cada chave é única e está associada a um valor. os mapas são usados para armazenar dados de forma eficiente e permitem acesso rápido aos valores com base nas chaves. em go, os mapas são implementados como tabelas hash, o que significa que eles oferecem desempenho constante para operações de inserção, exclusão e busca.