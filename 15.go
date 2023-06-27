package main

import (
	"errors"
	"fmt"
)

func contemNumero(numero int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("O slice está vazio")
	}

	for _, num := range slice {
		if num == numero {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	numero := 7
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	estaPresente, err := contemNumero(numero, slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("O número %d está presente no slice: %v \n", numero, estaPresente)
	}

	sliceVazio := []int{}
	estaPresenteVazio, errVazio := contemNumero(numero, sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Printf("O número %d está presente no slice: %v \n ", numero, estaPresenteVazio)
	}
}
