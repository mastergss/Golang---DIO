// Desafio: testar as funções de subtração e divisão;
// Utilizar convenção padrão para testes
// Aplicar triple 'A'
// Executar esteira de testes

package main

import "testing"

func ShouldSubCorrect (t *testing.T) { //func TestSubtrair(t *testing.T) {
	sub := subtrair(10, 2)
	resultado := 8

	if sub != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", sub)
	}
}

func ShouldSubIncorrect(t *testing.T) { //func TestSubtrair2(t *testing.T) {
	sub := subtrair(10, 4)
	resultado := 5

	if sub != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", sub)
	}
}

func ShouldDivCorrect(t *testing.T){ //func TestDividir(t *testing.T) {
	div := dividir(4, 2)
	resultado := 1

	if div != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", div)
	}
}

func ShouldDivIncorrect(t *testing.T) { //func TestDividir2(t *testing.T) {
	div := dividir(8, 4)
	resultado := 1

	if div != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", div)
	}
}