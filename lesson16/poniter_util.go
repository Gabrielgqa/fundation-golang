package main

func soma(a, b *int) int {
	*a = 50 // Não passando como ponteiro isso seria só uma cópia
	return *a + *b
}

func main() {
	n1 := 10
	n2 := 20
	soma(&n1, &n2)
	println(n1)
}
