package main

import(
    "fmt"
)

func main() {
    var n float64
    fmt.Scanf("%f",&n)
    i := int(n)
    fmt.Printf("%d",i%10)
}