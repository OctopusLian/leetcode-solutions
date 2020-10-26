func smallerNumbersThanCurrent(nums []int) []int {
	//解法1，暴力
    var ans []int
    for _, v := range nums {
        cnt := 0
        for _, w := range nums {
            if w < v {
                cnt++
            }
        }
        ans = append(ans, cnt)
    }
    return ans
}