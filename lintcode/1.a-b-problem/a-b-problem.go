/*
 * @Description:
 * @Author: neozhang
 * @Date: 2021-11-14 20:44:52
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-11 21:51:11
 */
package main

import "fmt"

func main() {
    fmt.Println(aplusb(1,2))
}

/**
 * @param a: An integer
 * @param b: An integer
 * @return: The sum of a and b
 */
func aplusb(a int, b int) int {
	return a + b
}