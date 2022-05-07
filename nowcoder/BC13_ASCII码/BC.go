package main

import(
    "fmt"
)

func main() {
    nums := []int{73,32,99,97,110,32,100,111,32,105,116,33}
    for _,num := range nums {
        //int转string
        fmt.Print(string(num))  //注意，不换行
    }
}