package main

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) recebeu(valor int) int {
	c.saldo += valor
	return c.saldo
}

func main() {
	gabriel := NewConta()
	gabriel.saldo = 200
	gabriel.recebeu(50)
	println(gabriel.saldo)
}
