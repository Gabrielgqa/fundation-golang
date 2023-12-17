package main

type Number interface {
	int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{"Gabriel": 200, "Ianca": 300}
	m2 := map[string]float64{"Gabriel": 200.30, "Ianca": 300.45}
	println(Soma(m))
	println(Soma(m2))
}
