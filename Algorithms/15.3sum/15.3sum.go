func threeSum(nums []int) [][]int {
    results := [][]int{}
	n := len(nums)
	if n == 0 || n < 3 {
		return results
	}
	sort.Ints(nums)  //首先，对数组进行排序
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {  //如果相邻两个数相等
			continue
		}
		target := -nums[i]
		left := i + 1  //左指针
		right := n - 1  //右指针
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				results = append(results, []int{nums[left], nums[right], nums[i]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}

	}
	return results
}
