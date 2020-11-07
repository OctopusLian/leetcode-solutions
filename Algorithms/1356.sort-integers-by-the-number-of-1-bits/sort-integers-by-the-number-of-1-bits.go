func onesCount(x int) (c int) {
    for ; x > 0; x /= 2 {
        c += x % 2
    }
    return
}

func sortByBits(a []int) []int {
    sort.Slice(a, func(i, j int) bool {
        x, y := a[i], a[j]
        cx, cy := onesCount(x), onesCount(y)
        return cx < cy || cx == cy && x < y
    })
    return a
}