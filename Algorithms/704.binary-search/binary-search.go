package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1 //初始化左右指针
	for left <= right {
		mid := (left + right) / 2 //定中间指针
		if nums[mid] > target {   //如果中间指针下标所对应的值大于target
			right-- //右指针向左移动一位
		} else if nums[mid] < target { //如果中间指针下标所对应的值小于target
			left++ //左指针向右移动一位
		} else if nums[mid] == target { //如果中间指针下标所对应的值等于target
			return mid //返回mid
		}
	}

	return -1 //否则返回-1
}

func main() {

}
