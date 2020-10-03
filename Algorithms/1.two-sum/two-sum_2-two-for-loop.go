func twoSum(nums []int, target int) []int {
	//暴力法，双循环
    for i, x := range nums {  //第一个循环
        for j := i + 1; j < len(nums); j++ {  //第二个循环
            if x+nums[j] == target {  //如果数组前一个值+后一个值 = target
                return []int{i, j}  //返回这两个值
            }
        }
    }
    return nil
}