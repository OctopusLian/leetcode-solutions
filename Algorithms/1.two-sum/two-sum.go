func twoSum(nums []int, target int) []int {
    var result = [2]int {0,0}
    if len(nums) < 2 {
        return nil
    }
    
    for i := 0 ; i < len(nums) - 1; i++ {
        for j := i + 1; j < len(nums); j++ {  
            if(nums[i] + nums[j] == target){
                result[0] = i
                result[1] = j
                return result[:]  //返回结果
            }
        }
    }
    return nil    
}

