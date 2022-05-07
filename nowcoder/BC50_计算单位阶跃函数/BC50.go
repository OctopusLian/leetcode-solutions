package main
import "fmt"
func main() {
    for {
        var a int
        _, err := fmt.Scanf("%d",&a)
        if err != nil{
            break
        }
        if a >0 {
            fmt.Println("1")
        } else if a == 0 {
            fmt.Println("0.5")
        } else {
            fmt.Println("0")
        }
    }
     
     
}
