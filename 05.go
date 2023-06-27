package main

import "fmt"

func encontrarPosicao(slice []int, valor int) int {
	for i, v := range slice {
		if v == valor {
			return i
		}
	}

	return -1
}

func main() {
	numeros := []int{10, 20, 30, 40, 50}
	valor := 30
	posicao := encontrarPosicao(numeros, valor)
	fmt.Printf("O valor %d está na posição %d \n", valor, posicao)

	valorNaoEncontrado := 60
	posicaoNaoEncontrada := encontrarPosicao(numeros, valorNaoEncontrado)
	fmt.Printf("O valor %d não foi encdontrado. Retornou: %d \n", valorNaoEncontrado, posicaoNaoEncontrada)
}
