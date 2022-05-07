package main
 
import (
    "fmt"
    "math"
)
 
func main(){
    var a,b float64
    fmt.Scanf("%f %f",&a,&b)
    fmt.Printf("%.2f",a/math.Pow(b/100,2))
}