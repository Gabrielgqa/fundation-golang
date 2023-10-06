package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0]) // Dropa os elementos a partir do 1º índice
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4]) // Dropa os elementos a partir do 4º índice
	fmt.Printf("len=%d cap=%d %v\n", len(s[4:]), cap(s[4:]), s[4:]) // Dropa os elementos a antes do 4º índice (posição dos : (dois pontos)) a capacidade tbm diminui

	// Ao adicionar itens ao slice a capacidade atual dobra. Sempre pensar em como manejar para não usar tantos espaços
	s = append(s, 12)
	s = append(s, 14)
	s = append(s, 16)
	s = append(s, 18)
	s = append(s, 20)
	s = append(s, 22)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
