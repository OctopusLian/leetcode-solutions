func twoSum(nums []int, target int) []int {
    left,right := 0,len(nums)-1
    for ;left < right; {
            if (nums[left] + nums[right] == target) {
                return []int{nums[left],nums[right]}
            }

            if nums[left]>target-nums[right]{
                //right过大，向左移动一位
                right--
                continue
            }else{
                //left过大，像右移动一位
                left++
                continue
            }
    }
    return nil
}
