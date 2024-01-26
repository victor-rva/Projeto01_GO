package main

import "honnef.co/go/tools/printf"

type Car struct {
	Model string
	Color string
}

// metodo
func (c Car) Start() {
	println(c.Model + "is started")
}

// funcao
func soma(x int, y int) int {
	return x + y
}

func main() {
	car := Car{ // declarando e atribuindo variavel car, serve para não precisar declarar antes a variavel
		Model: "Ferrari",
		Color: "Red",
	}
	car.Model = "Fiat" // alterando o valor do atributo Model
	println(car.Model)
}