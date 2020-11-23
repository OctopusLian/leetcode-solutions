func findMinArrowShots(points [][]int) int {
	//解法1，排序+贪心
    if len(points) == 0 {
        return 0
    }
    sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
    maxRight := points[0][1]
    ans := 1
    for _, p := range points {
        if p[0] > maxRight {
            maxRight = p[1]
            ans++
        }
    }
    return ans
}