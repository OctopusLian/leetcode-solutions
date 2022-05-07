package main
 
import "fmt"
 
func main() {
    var a,b,sum int
    fmt.Scanf("0x%x 0%o",&a,&b)
    sum = a + b
    fmt.Printf("%d\n",sum)
}