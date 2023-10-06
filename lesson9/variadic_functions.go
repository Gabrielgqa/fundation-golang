package main

import "fmt"

func main() {
	fmt.Println(soma(50, 6, 5, 3, 8))
}

// Multiplos retornos com diferentes tipos
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
