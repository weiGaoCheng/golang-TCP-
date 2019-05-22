package main

import "fmt"

//----------------------------抽象层
type Car2 interface {
	Run()
}
type Driver2 interface {
	Driver2(car Car)
}

//----------------------------实现层
//BMW
type BMW struct {
	Car2
}
func (bmw *BMW)Run(){
	fmt.Println("BMW is running...")
}
//BenZ
type Benz struct {
	Car2
}
func (benz *Benz)Run(){
	fmt.Println("BenZ is running...")
}
//Driver

//----------------------------业务逻辑层
