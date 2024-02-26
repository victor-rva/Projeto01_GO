package main

import (
	"fmt"
	"time"
)

func processando() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main(){
	canal := make(chan int)
	go func(){
		for i := 0; i < 10; i++{
			canal <- 1
			fmt.Println("Jogou no canal", i)
		}	
	}()

	go worker(canal, 1)
	worker(canal, 2)
}

func worker(canal chan int, workerID int){
	for{
		fmt.Println("Recebeu do canal", <- canal, "no worker", workerID)
		time.Sleep(time.Second)
	}
}

// //Thread 1
// func main() {
// 	go processando() // thread 2
// 	go processando() // thread 3
// 	processando() // thread 1
// }



// func main() {
// 	order, err := entity.NewOrder("1",10,1)
// 	if err != nil{
// 		println(err.Error())
// 	} else{
// 		println(order.ID)
// 	}
// }

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

// func (c Car) ChangeColor(color string){
// 	c.Color = color // duplicando o valor de c.color na memoria
// 	println("New color: " + c.Color)
// }

type Car struct {
	Model string
	Color string
}

// metodo
func (c Car) Start() {
	println(c.Model + "has been started")
}



func (c *Car) ChangeColor(color string) { // o * serve como ponteiro para indicar que a mudança do valor vai ser direto da memória we não uma cópia
	c.Color = color
	println("New color: " + c.Color)
}

// funcao
func soma(x int, y int) int{
	return x + y
}
