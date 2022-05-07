package main

import(
    "fmt"
)

func main() {
    var seconds int
    fmt.Scan(&seconds)
    fmt.Printf("%d %d %d", seconds/3600, seconds%3600/60, seconds%60)
}