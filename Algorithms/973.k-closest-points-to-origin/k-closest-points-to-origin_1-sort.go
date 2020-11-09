func kClosest(points [][]int, K int) [][]int {
	//解法1,排序
    sort.Slice(points, func(i, j int) bool {
        p, q := points[i], points[j]
        return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
    })
    return points[:K]
}