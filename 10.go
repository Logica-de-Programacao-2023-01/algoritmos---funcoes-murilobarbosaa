package main

import (
	"errors"
	"fmt"
	"sort"
)

func ordenarValores(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	valoresOrdenados := make([]int, len(slice))
	copy(valoresOrdenados, slice)
	sort.Ints(valoresOrdenados)

	return valoresOrdenados, nil
}

func main() {
	valores := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	valoresOrdenados, err := ordenarValores(valores)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Valores ordenados:", valoresOrdenados)
	}

	valoresVazios := []int{}
	valoresOrdenadosVazios, errVazios := ordenarValores(valoresVazios)
	if errVazios != nil {
		fmt.Println("Erro:", errVazios)
	} else {
		fmt.Println("Valores ordenados:", valoresOrdenadosVazios)
	}
}
