package main

import (
	"errors"
	"fmt"
)

func aplicarFuncao(slice []int, funcao func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	resultado := make([]int, len(slice))
	for i, valor := range slice {
		resultado[i] = funcao(valor)
	}

	return resultado, nil
}

func dobrarvalor(numero int) int {
	return numero * 2
}

func main() {
	valores := []int{1, 2, 3, 4, 5}
	resultado, err := aplicarFuncao(valores, dobrarvalor)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}
