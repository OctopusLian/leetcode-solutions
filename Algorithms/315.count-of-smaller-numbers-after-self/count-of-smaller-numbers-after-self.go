func countSmaller(nums []int) []int {
    n := len(nums)
    
    var res []int
    for i:=0;i < n;i++{
         tem := 0
        for j :=i;j < n;j ++{
            if nums[j] < nums[i]{
                tem++
               
            }
        }
         res  = append(res,tem)
    }
    
    return res
}