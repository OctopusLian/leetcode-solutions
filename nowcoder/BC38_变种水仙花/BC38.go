package main
 
import "fmt"
 
func main() {
    var a int
    sum := 0
    var p1,p2 int
    for a = 10000;a < 100000;a++{
        for i := 1;i < 5;i++){
            p1 = a/(int)pow(10,i)
            p2 = a%(int)pow(10,i)
            sum += p1*p2
        }
        if (a == sum) {
             fmt.Printf("%d ",sum)
        }
        sum = 0
    }
}
