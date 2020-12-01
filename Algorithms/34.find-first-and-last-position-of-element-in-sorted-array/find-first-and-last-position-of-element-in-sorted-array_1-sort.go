func searchRange(nums []int, target int) []int {
    l, r := sort.SearchInts(nums, target), sort.SearchInts(nums, target+1)
  
    if l < 0 || l == len(nums) || nums[l] != target {
        return []int{-1, -1}
    } 

    return []int{l, r-1}
}