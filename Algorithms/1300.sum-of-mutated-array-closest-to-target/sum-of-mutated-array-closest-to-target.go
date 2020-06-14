func findBestValue(arr []int, target int) int {
	//解法1，枚举+二分查找
    sort.Ints(arr)
    n := len(arr)
    prefix := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        prefix[i] = prefix[i-1] + arr[i-1]
    }
    r := arr[n-1]
    ans, diff := 0, target
    for i := 1; i <= r; i++ {
        index := sort.SearchInts(arr, i)
        if index < 0 {
            index = -index - 1
        }
        cur := prefix[index] + (n - index) * i
        if abs(cur - target) < diff {
            ans = i
            diff = abs(cur - target)
        }
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -1 * x
    }
    return x
}