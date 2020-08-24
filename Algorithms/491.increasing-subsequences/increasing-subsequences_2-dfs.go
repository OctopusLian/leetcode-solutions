func findSubsequences(nums []int) [][]int {
	//解法2，回溯DFS
    var res [][]int
    var dfs func(int, []int) 
    dfs = func (index int, lis []int) {
        if len(lis) >= 2 {
            dest := make([]int, len(lis))
            copy(dest, lis)
            res = append(res, dest)
        }
        //保证一层循环中不会选择重复的值就ok了，画一下可能会比较清晰
        var visit [201]bool
        for i := index; i < len(nums); i++ {
            if visit[nums[i]+100] {
                continue
            }
            if len(lis) == 0 || nums[i] >= lis[len(lis)-1] {
                visit[nums[i]+100] = true
                dfs(i+1, append(lis, nums[i]))
            }
        }
    }
    dfs(0, []int{})
    return res
}