package main

import(
    "fmt"
)

func main() {
    var i,a int
    fmt.Scan(&a)  //输入
    for i = 0;a != 1;i++ {
        if a%2 == 0{
            a=a/2
        }else {
            a=(3*a+1)/2
        }
        
    }
    fmt.Println(i)
}