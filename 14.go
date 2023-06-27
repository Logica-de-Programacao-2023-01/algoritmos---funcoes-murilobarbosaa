package main

import (
	"errors"
	"fmt"
)

func intersecao(slice1 []int, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("Um dos slices está vazio")
	}

	intersecao := make([]int, 0)
	mapaSlice1 := make(map[int]bool)

	for _, num := range slice1 {
		mapaSlice1[num] = true
	}

	for _, num := range slice2 {
		if mapaSlice1[num] {
			intersecao = append(intersecao, num)
		}
	}

	return intersecao, nil
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}
	resultado, err := intersecao(slice1, slice2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Interseção:", resultado)
	}

	sliceVazio := []int{}
	resultadoVazio, errVazio := intersecao(slice1, sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Interseção", resultadoVazio)
	}

}
