package main

import(
    "fmt"
)

func main() {
    for {
        var n int
        _,err:=fmt.Scan(&n)
        if err!=nil{
            break
        }
        res := 1
        for i := 0;i < n;i++ {
            res <<= 1
        }
        fmt.Println(res)
    }
}