//Pacote container/list fornece uma implementação de lista duplamente ligada. Ele inclui funções para criar e manipular listas, como adicionar, remover e iterar sobre os elementos da lista. O pacote container/list é útil para casos em que a ordem dos elementos é importante e quando é necessário inserir ou remover elementos em qualquer posição da lista de forma eficiente.

//Exemplo de uso do pacote container/list:

package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Criando uma nova lista e adicionando elementos a ela
	l := list.New()
	
	// Adicionando elementos à lista usando PushBack, PushFront, InsertBefore e InsertAfter
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)


	// Iterando através da lista e imprimindo seu conteúdo.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}