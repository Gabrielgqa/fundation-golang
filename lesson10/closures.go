package main

import "fmt"

func main() {
	total := func() int {
		return soma(50, 6, 5, 3, 8) * 3
	}()

	fmt.Println(total)
}

// Multiplos retornos com diferentes tipos
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
