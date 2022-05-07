package main

import (
	"fmt"
)

func main() {
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				sum1 := a*100 + b*10 + c
				sum2 := b*100 + c*10 + c
				if sum1 >= 100 && sum2 >= 100 && (sum1+sum2) == 532 {
					fmt.Printf("%d %d %d\n", a, b, c)
				}
			}
		}
	}

	return
}
