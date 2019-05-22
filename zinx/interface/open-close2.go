package main

import "fmt"

//一个抽象的业务员
type AbstractBacker interface {
	Dobusi()
}
//支付业务
type PayBacker struct {

}
//框架层
func BackerBusiness (backer AbstractBacker){
	fmt.Println("实现了业务:")
	backer.Dobusi()
}

