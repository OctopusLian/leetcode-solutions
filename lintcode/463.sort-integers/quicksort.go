/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-11 22:15:00
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-11 22:15:01
 */
package main

func main() {

}

func sortIntegers (A *[]int)  {
    // write your code here
    if A == nil || len(A) == 0 {
        return
    }
    QuickSort(A,0,len(A)-1)
}

func QuickSort(A *[]int,start int,end int) {
    if start >= end {
		return
	}
}