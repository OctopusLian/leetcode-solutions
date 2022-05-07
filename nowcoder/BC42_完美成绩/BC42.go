package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    if n >= 90 && n <= 100 {
        fmt.Printf("Perfect\n")
    }
}