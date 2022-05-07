package main
  
import "fmt"
  
func main() {
    for {
        var a,b int
        _,err := fmt.Scanf("%d %d\n",&a,&b)
        if err != nil {
            break
        }
        if a > b {
            fmt.Printf("%d>%d\n",a,b)
        } else if a == b {
            fmt.Printf("%d=%d\n",a,b)
        } else {
            fmt.Printf("%d<%d\n",a,b)
        }
    }
}
