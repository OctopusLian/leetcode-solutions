package main
 
import (
    "fmt"
    "math"
)
 
func main() {
    var a, b, c float64
    fmt.Scanf("%f %f %f", &a, &b, &c)
    sum := a + b + c
    /*海伦公式p=(a+b+c)/2,s=sqrt(p(p-a)(p-b)(p-c));*/
    p := sum / 2
    area := math.Sqrt(p * (p - a) * (p - b) * (p - c))
    fmt.Printf("circumference=%.2f area=%.2f", sum, area)
}