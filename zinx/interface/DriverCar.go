package main

import "fmt"

// --- 抽象层 ---
type Car interface {
	Run()
}
type Driver interface {
	Drive(car Car)
}

// --- 实现层 ---
type BenZ struct {
	//...
}

func (benz *BenZ) Run() {
	fmt.Println("Benz is running...")
}

type Bmw struct {
	//...
}

func (bmw *Bmw) Run() {
	fmt.Println("Bmw is running...")
}


type Z3 struct {
	//...
}

func (z3 *Z3) Drive(car Car) {
	fmt.Println("zhan3 driver car")
	car.Run()
}


type L4 struct {
	//...
}

func (l4 *L4) Drive(car Car) {
	fmt.Println("li4 driver car")
	car.Run()
}


//业务逻辑层
func main() {
	//业务1 张3 开宝马
	var bmw Car
	bmw = &Bmw{}

	var zhang3 Z3
	zhang3.Drive(bmw)



}
