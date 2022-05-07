package main
import "fmt"
func main(){
    var a,b int
    fmt.Scanf("%d %d",&a,&b)
    fmt.Printf("%d",(a+b)%100)  //两个数加起来再取100的余
}