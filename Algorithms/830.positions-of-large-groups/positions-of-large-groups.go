func largeGroupPositions(s string) [][]int {
	//解法1，一次遍历
    var ans [][]int
    cnt := 1
    for i := range s {
        if i == len(s)-1 || s[i] != s[i+1] {
            if cnt >= 3 {
                ans = append(ans, []int{i - cnt + 1, i})
            }
            cnt = 1
        } else {
            cnt++
        }
    }
    return ans
}