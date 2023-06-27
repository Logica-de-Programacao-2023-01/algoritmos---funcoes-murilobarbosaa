package main

import (
	"errors"
	"fmt"
	"math"
)

func isPrimo(numero int) bool {
	if numero < 2 {
		return false
	}

	limite := int(math.Sqrt(float64(numero)))

	for i := 2; i <= limite; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

func numerosPrimos(numero int) ([]int, error) {
	if numero < 0 {
		return nil, errors.New("O número é negativo")
	}

	primos := []int{}

	for i := 2; i <= numero; i++ {
		if isPrimo(i) {
			primos = append(primos, i)
		}
	}

	return primos, nil
}

func main() {
	numero := 20
	primos, err := numerosPrimos(numero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Número: %d\n", numero)
		fmt.Printf("Números primos menores ou iguais: %v\n", primos)
	}

	numeroNegativo := -10
	primosNegativo, errNegativo := numerosPrimos(numeroNegativo)
	if errNegativo != nil {
		fmt.Println("Erro:", errNegativo)
	} else {
		fmt.Printf("Número: %d\n", numeroNegativo)
		fmt.Printf("Números primos menores ou iguais: %v\n", primosNegativo)
	}
}
