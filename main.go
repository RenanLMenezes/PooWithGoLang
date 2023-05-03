package main

import "fmt"

type ContaCorrente struct {
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}

}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor deve ser maior que 0", c.saldo
	}

}

func main() {
	cliente1 := ContaCorrente{}
	cliente1.titular = "Renan"
	cliente1.saldo = 500

	fmt.Println(cliente1.saldo)
	fmt.Println(cliente1.Sacar(-600))
	fmt.Println(cliente1.Sacar(100))
	fmt.Println(cliente1.saldo)
	fmt.Println(cliente1.Depositar(-100))
	fmt.Println(cliente1.Depositar(100))

}
