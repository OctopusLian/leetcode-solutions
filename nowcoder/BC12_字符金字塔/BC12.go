package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)
	for i := 0; i < 5; i++ { //金字塔总共有5层
		for j := 0; j < 4-i; j++ {
			fmt.Print(" ") //输出空格
		}
		for k := 0; k < i+1; k++ {
			fmt.Printf("%s ", input) //输出字符
		}
		fmt.Printf("\n") //每一层循环结束后要换行
	}
}

/*
# go run BC12.go
1
    1
   1 1
  1 1 1
 1 1 1 1
1 1 1 1 1
*/
