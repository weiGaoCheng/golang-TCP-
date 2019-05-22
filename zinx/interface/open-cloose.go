
package main

import "fmt"

//抽象的业务员
type AbstractBanker interface {
	DoBusi() //抽象的接口 ，业务接口
}

//存款的业务员
type SaveBanker struct {
	AbstractBanker
}

func (sb *SaveBanker) DoBusi () {
	fmt.Println("进行的存款")
}

//转账的业务员
type TransBanker struct {
	AbstractBanker
}

func (sb *TransBanker) DoBusi () {
	fmt.Println("进行的转账")
}

//支付的业务员
type PayBanker struct {
	AbstractBanker
}

func (pb *PayBanker) DoBusi() {
	fmt.Println("进行的支付")
}



//实现一个架构层(基于抽象的接口来进行封装 )
func BankerBusiness(banker AbstractBanker) {
	fmt.Println("进行了业务")
	banker.DoBusi() //多态的现象
}




func main() {
	/*
	sb := &SaveBanker{}
	sb.DoBusi()

	tb := &TransBanker{}
	tb.DoBusi()
	*/
	BankerBusiness(&SaveBanker{})
	BankerBusiness(&TransBanker{})
	BankerBusiness(&PayBanker{})
}