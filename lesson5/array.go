package main

import "fmt"

func main() {
	var meuArray [3]int

	meuArray[0] = 1
	meuArray[1] = 5
	meuArray[2] = 7

	for i, v := range meuArray {
		fmt.Printf("Valor do indice %d é %d \n", i, v)
	}

	fmt.Println(len(meuArray))
	fmt.Println(meuArray[len(meuArray)-1])

}
