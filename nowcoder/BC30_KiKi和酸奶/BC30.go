package main

import(
    "fmt"
)

func main() {
    var n,h,m int
    fmt.Scanf("%d %d %d",&n,&h,&m)
    if m % h == 0 {
        //刚好喝完 m/h 瓶的酸奶
        fmt.Println(n - m / h)
    } else {
        //喝完 m/h 瓶酸奶后，正在喝第 ( m / h - 1) 瓶酸奶
        fmt.Println(n - m / h - 1)
    }
}