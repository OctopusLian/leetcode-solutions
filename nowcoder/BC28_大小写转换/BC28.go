package main
 
import (
    "bufio"
    "fmt"
    "os"
)
 
func main() {
    sc := bufio.NewScanner(os.Stdin)
    for sc.Scan() {
        bytes := sc.Bytes()
        fmt.Printf("%c\n", bytes[0] + 32)  //大写转小写
    }
}