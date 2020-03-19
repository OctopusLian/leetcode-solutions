package main

import (
	"fmt"
)

type Plane struct {
	Signal int
	Msg    string
	X      int
	Y      int
	Z      int
	Status string
}

func (p *Plane) Distinguish(num int) (status int, signal int) {

}

func ReadFile(path string) {

}

func main() {
	//1，读取文本文件
	ReadFile("./data.txt")

	//2，读取输入

	//3，做判断，并输出

}
