package main

import "fmt"

func calcularMedia(valores []int) float64 {
	soma := 0

	for _, valor := range valores {
		soma += valor
	}

	media := float64(soma) / float64(len(valores))
	return media
}

func main() {
	numeros := []int{10, 20, 30, 40, 50}
	media := calcularMedia(numeros)
	fmt.Println("Média:", media)
}
