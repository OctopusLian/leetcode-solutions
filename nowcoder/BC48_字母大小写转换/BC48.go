package main

import "fmt"

func main() {
    for {
        var s string
        _, err := fmt.Scan(&s)
        if err != nil {
            break
        }
        var result string
        if s[0] >= 'A' && s[0] <= 'Z' {
            result = string(s[0] + 32)
        } else if s[0] >= 'a' && s[0] <= 'z' {
            result = string(s[0] - 32)
        }
        fmt.Println(result)
    }
}