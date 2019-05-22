package main

import "fmt"

/*
模拟组装2台电脑
--- 抽象层 ---
有显卡Card  方法display
有内存Memory 方法storage
有处理器CPU   方法calculate

--- 实现层 ---
有 Intel因特尔公司 、产品有(显卡、内存、CPU)
有 Kingston 公司， 产品有(内存)
有 NVIDIA 公司， 产品有(显卡)

--- 逻辑层 ---
1. 组装一台Intel系列的电脑，并运行
2. 组装一台 Intel CPU  Kingston内存 NVIDIA显卡的电脑，并运行
*/

// ---- 抽象层 ----
type Card interface {
	Display()
}

type Memory interface {
	Storage()
}

type CPU interface {
	Calculate()
}

type Computer struct {
	cpu CPU
	mem Memory
	card Card
}

//初始化Computer对象的方法
func NewComputer(cpu CPU, mem Memory, card Card) *Computer {
	//com 对象
	return 	&Computer{
		cpu:cpu,
		mem:mem,
		card:card,
	}
}

func (this *Computer) Work() {
	//多态的现象
	this.cpu.Calculate()
	this.card.Display()
	this.mem.Storage()
}


//0---0 实现层
//IntelCPU
type IntelCPU struct {
	CPU
}
func (this *IntelCPU)Calculate() {
	fmt.Println("Intel cpu 开始计算了...")
}

//IntelCard
type IntelCard struct {
	Card
}
func (this *IntelCard) Display() {
	fmt.Println("Intel card 开始显示了...")
}
//IntelMem
type IntelMem struct {
	Memory
}
func (this *IntelMem) Storage() {
	fmt.Println("Intel mem 开始存储了...")
}
//kinston的内存
type KinstonMem struct {
	Memory
}

func (km *KinstonMem) Storage() {
	fmt.Println("Kingston mem 开始存储了")
}
//NVIDIA的显卡
type NCard struct {
	Card
}

func (nc *NCard) Display() {
	fmt.Println("n card 开始显示了")
}


func main() {
	//1 组装一台intel系列的电脑
	com1 := NewComputer(&IntelCPU{}, &IntelMem{}, &IntelCard{})
	com1.Work()

	fmt.Println("------")
	//2 组装第二天电脑
	com2 := NewComputer(&IntelCPU{}, &KinstonMem{}, &NCard{})
	com2.Work()
}


