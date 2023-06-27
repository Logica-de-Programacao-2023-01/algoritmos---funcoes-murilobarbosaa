package main

import (
	"errors"
	"fmt"
)

func numerosPares(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	pares := make([]int, 0)
	for _, num := range slice {
		if num%2 == 0 {
			pares = append(pares, num)
		}
	}

	return pares, nil
}

func main() {
	valores := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pares, err := numerosPares(valores)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números pares:", pares)
	}

	valoresVazios := []int{}
	paresVazios, errVazios := numerosPares(valoresVazios)
	if errVazios != nil {
		fmt.Println("Erro", errVazios)
	} else {
		fmt.Println("Números pares:", paresVazios)
	}
}
