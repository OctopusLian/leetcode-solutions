package main

import(
    "fmt"
)

func main() {
    var str string
    fmt.Scan(&str)
    var result string
    strLen := len(str)
    for i := 0; i < strLen; i++ {
        result = result + fmt.Sprintf("%c", str[strLen-i-1])
    }
    fmt.Print(result)
}