func less(p, q []int) bool {
    return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
}

func kClosest(points [][]int, k int) (ans [][]int) {
	//解法2，快速选择
    rand.Shuffle(len(points), func(i, j int) {
        points[i], points[j] = points[j], points[i]
    })

    var quickSelect func(left, right int)
    quickSelect = func(left, right int) {
        if left == right {
            return
        }
        pivot := points[right] // 取当前区间 [left,right] 最右侧元素作为切分参照
        lessCount := left
        for i := left; i < right; i++ {
            if less(points[i], pivot) {
                points[i], points[lessCount] = points[lessCount], points[i]
                lessCount++
            }
        }
        // 循环结束后，有 lessCount 个元素比 pivot 小
        // 把 pivot 交换到 points[lessCount] 的位置
        // 交换之后，points[lessCount] 左侧的元素均小于 pivot，points[lessCount] 右侧的元素均不小于 pivot
        points[right], points[lessCount] = points[lessCount], points[right]
        if lessCount+1 == k {
            return
        } else if lessCount+1 < k {
            quickSelect(lessCount+1, right)
        } else {
            quickSelect(left, lessCount-1)
        }
    }
    quickSelect(0, len(points)-1)
    return points[:k]
}