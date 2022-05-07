package main

import(
    "fmt"
)

func main() {
    var a,res int
    fmt.Scanf("%d",&a)
    //qian := a / 1000  //千位
    //bai := (a - qian*1000) / 100  //百位
    //shi := (a - qian*1000 - bai*100) / 10  //十位
    //ge := a - qian*1000 - bai*100 - shi*10  //个位
    //res := ge*1000 + bai*10 + shi*100 + qian
    for a > 0 {
        res = a % 10
        a = a / 10
        fmt.Print(res)  //注意不要换行
    }
}