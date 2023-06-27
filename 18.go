package main

import (
	"errors"
	"fmt"
)

func aplicarFuncao(slice []int, funcao func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("O slice está vazio")
	}

	resultado := 0
	for _, valor := range slice {
		resultado += funcao(valor)
	}

	return resultado, nil
}

func dobrar(numero int) int {
	return numero * 2
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	somaDobro, err := aplicarFuncao(slice, dobrar)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Slice: %v\n", slice)
		fmt.Printf("Soma do dobro: %d\n", somaDobro)
	}

	sliceVazio := []int{}
	somaDobroVazio, errVazio := aplicarFuncao(sliceVazio, dobrar)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Printf("Slice: %v\n", sliceVazio)
		fmt.Printf("Soma do dobro: %d\n", somaDobroVazio)
	}
}
