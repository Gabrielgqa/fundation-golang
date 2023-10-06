package main

type Conta struct {
	saldo int
}

func (c *Conta) recebeu(valor int) int {
	c.saldo += valor
	return c.saldo
}

func main() {
	gabriel := Conta{
		saldo: 100,
	}
	gabriel.recebeu(50)
	println(gabriel.saldo)
}
