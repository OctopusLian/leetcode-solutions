func twoSum(numbers []int, target int) []int {
	//解法1，双指针法
    low, high := 0, len(numbers) - 1  //双指针，一个指向头，另一个指向尾
    for low < high {
        sum := numbers[low] + numbers[high]
        if sum == target {
            return []int{low + 1, high + 1}
        } else if sum < target {
            low++  //sum小于目标值，左指针向右移动一个元素
        } else {
            high--  //sum大于目标值，右指针向左移动一个元素
        }
    }
    return []int{-1, -1}
}