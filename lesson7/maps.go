package main

import "fmt"

func main() {
	salarios := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	fmt.Println(salarios["Wesley"])
	delete(salarios, "Wesley")
	salarios["Gabriel"] = 5000
	fmt.Println(salarios)

	// Duas formas de criar maps sem valor inicial
	sal := make(map[string]int)
	sal1 := map[string]int{}

	fmt.Println(sal, sal1)

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d", nome, salario)
	}
	// Blank identifier é ignorar um valor colocando _ (underline)
	for _, salario := range salarios {
		fmt.Printf("O salario é %d", salario)
	}
}
