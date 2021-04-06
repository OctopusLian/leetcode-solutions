package main

import(
	"fmt"
)

func main() {
	numss := []int{2,5,3,5}
	fmt.Print(purchasePlans(numss,6))
}

func purchasePlans(nums []int, target int) int {
    count := 0
    
    for short := 0;short < len(nums);short++ {
        for fast := 1;fast < len(nums);fast++ {
            if nums[short] + nums[fast] <= target {
                count = count + 1
            }
        }
    }
    
    return count
}