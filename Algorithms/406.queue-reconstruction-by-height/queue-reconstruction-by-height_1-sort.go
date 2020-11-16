func reconstructQueue(people [][]int) [][]int {
	//解法1，排序
    sort.Slice(people, func(i, j int) bool {
        a, b := people[i], people[j]
        return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
    })
    ans := make([][]int, len(people))
    for _, person := range people {
        spaces := person[1] + 1
        for i := range ans {
            if ans[i] == nil {
                spaces--
                if spaces == 0 {
                    ans[i] = person
                    break
                }
            }
        }
    }
    return ans
}