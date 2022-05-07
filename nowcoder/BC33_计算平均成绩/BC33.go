package main
 
import "fmt"
 
func main() {
    var a,b,c,d,e int
    var sum,avg float64
    fmt.Scanf("%d%d%d%d%d",&a,&b,&c,&d,&e)
    sum = float64(a + b + c + d + e)
    avg = sum / 5
    fmt.Printf("%.1f",avg)
}