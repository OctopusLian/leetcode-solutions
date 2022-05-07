package main
 
import "fmt"
 
func main() {
    var height,weight,bmi float64
    fmt.Scanf("%f%f",&height,&weight)
    bmi = height / (weight*weight)
    if(bmi>18.5&&bmi<23.9){
        fmt.Printf("Normal")
    }else{
        fmt.Printf("Abnormal")
    }
}