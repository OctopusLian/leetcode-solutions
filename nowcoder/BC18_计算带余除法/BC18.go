package main

import(
    "fmt"
)

func main(){
    var a,b int
    _, _ = fmt.Scanf("%d %d", &a, &b)
    fmt.Println(a/b,a%b)  //输出商和余数

}