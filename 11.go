package main

import (
	"errors"
	"fmt"
	"math"
)

func isPrimo(numero int) (bool, error) {
	if numero < 0 {
		return false, errors.New("O número é negativo")
	}

	if numero < 2 {
		return false, nil
	}

	limite := int(math.Sqrt(float64(numero)))
	for i := 2; i <= limite; i++ {
		if numero%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	numero := 17
	primo, err := isPrimo(numero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("%d é primo: %v \n", numero, primo)
	}

	numeroNegativo := -10
	primoNegativo, errNegativo := isPrimo(numeroNegativo)
	if errNegativo != nil {
		fmt.Println("Erro:", errNegativo)
	} else {
		fmt.Printf("%d é primo: %v \n", numeroNegativo, primoNegativo)
	}
}
