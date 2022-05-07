package main
 
import (
    "fmt"
)

func main() {
    var a, b, c float32
    fmt.Scanf("%f %f %f", &a, &b, &c)
    sum := a + b + c
    fmt.Printf("%.2f %.2f", sum, sum / 3)
}