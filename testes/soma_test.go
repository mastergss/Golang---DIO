package main

import "testing"

// Convenção de nomes - a assinatura do método deve iniciar com Should(função)Correct/Incorrect

func ShouldSumCorrect(t *testing.T) { //func TestSoma(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func ShouldSumIncorrect(t *testing.T) { //func TestSoma2(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 7

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func ShouldMultCorrect(t *testing.T) { //func TestMultiplica(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 100

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor obtido: ", teste)
	}
}

func ShouldMultIncorrect(t *testing.T) { //func TestMultiplica2(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 2560

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor obtido: ", teste)
	}
}
