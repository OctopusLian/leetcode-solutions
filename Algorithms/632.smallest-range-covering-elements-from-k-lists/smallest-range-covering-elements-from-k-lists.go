func smallestRange(nums [][]int) []int {
	//解法1，哈希表+滑动窗口
    size := len(nums)
    indices := map[int][]int{}
    xMin, xMax := math.MaxInt32, math.MinInt32
    for i := 0; i < size; i++ {
        for _, x := range nums[i] {
            indices[x] = append(indices[x], i)
            xMin = min(xMin, x)
            xMax = max(xMax, x)
        }
    }
    freq := make([]int, size)
    inside := 0
    left, right := xMin, xMin - 1
    bestLeft, bestRight := xMin, xMax
    for right < xMax {
        right++
        if len(indices[right]) > 0 {
            for _, x := range indices[right] {
                freq[x]++
                if freq[x] == 1 {
                    inside++
                }
            }
            for inside == size {
                if right - left < bestRight - bestLeft {
                    bestLeft, bestRight = left, right
                }
                for _, x := range indices[left] {
                    freq[x]--
                    if freq[x] == 0 {
                        inside--
                    }
                }
                left++
            }
        }
    }
    return []int{bestLeft, bestRight}
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}