func containsDuplicate(nums []int) bool {
    if len(nums) == 0 || len(nums) == 1 {
        //考虑nums特殊情况
        return false
    }

    count := 0
    for i:=0;i<len(nums);i++ {
        for j := i+1;j<len(nums);j++ {
            if nums[i] == nums[j] {
                count++
            }
        }    
    }

    if count != 0 {
        return true
    } else {
        return false
    }
}
