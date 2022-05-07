package main

import(
    "fmt"
)

func main() {
    var input string
    fmt.Scanf("%s",&input)
    fmt.Printf("year=%v\nmonth=%v\ndate=%v\n",input[:4], input[4:6], input[6:])
}