package main

import(
    "fmt"
)

const pi = 3.1415926

func main() {
    var r,v float64
    fmt.Scan(&r)
    v = (4 * pi * r * r * r) / 3
    fmt.Printf("%.3f",v)
}