package main

import "fmt"

type ContaCorrente struct {
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
}

func main() {
	// var cliente1 ContaCorrente = ContaCorrente{}
	cliente1 := ContaCorrente{titular: "Renan", saldo: 100.5}
	cliente2 := ContaCorrente{"Jão", 0002, 123457, 0}
	fmt.Println(cliente1, cliente2)
}
