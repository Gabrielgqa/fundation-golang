package main

import "fmt"

func main() {
	fmt.Println(soma(50, 6))
}

// Multiplos retornos com diferentes tipos
func soma(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}
