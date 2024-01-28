package main

import "github.com/victor-rva/projeto01_GO/internal/entity"

//import "honnef.co/go/tools/printf"

type Car struct {
	Model string
	Color string
}

// metodo
func (c Car) Start() {
	println(c.Model + "has been started")
}

// func (c Car) ChangeColor(color string){
// 	c.Color = color // duplicando o valor de c.color na memoria
// 	println("New color: " + c.Color)
// }

func (c *Car) ChangeColor(color string){ // o * serve como ponteiro para indicar que a mudança do valor vai ser direto da memória we não uma cópia
	c.Color = color
	println("New color: " + c.Color)
}

// funcao
func soma(x int, y int) int {
	return x + y
}

func main() {
	order, err := entity.NewOrder("1",10,1)
	if err != nil{
		println(err.Error())
	} else{
		println(order.ID)
	}
}

// func main() {
// 	car := Car{ // declarando e atribuindo variavel car, serve para não precisar declarar antes a variavel
// 		Model: "Ferrari",
// 		Color: "Red",
// 	}
// 	car.Start()
// 	car.ChangeColor("Blue")
// 	println(car.Color)
// 	car.Model = "Fiat" // alterando o valor do atributo Model
// 	println(car.Model)
// }

//para pegar o endereço de memória se utiliza & do lado da variavel