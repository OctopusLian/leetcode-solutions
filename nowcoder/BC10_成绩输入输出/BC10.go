package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)                    //一次输入三个数,中间用空格隔开,例如 10 20 30
	fmt.Printf("score1=%d,score2=%d,score3=%d", a, b, c) //score1=10,score2=20,score3=30
}
