func topKFrequent(nums []int, k int) []int {
	//解法1，快速排序法
    occurrences := map[int]int{}
    for _, num := range nums {
        occurrences[num]++
    }
    values := [][]int{}
    for key, value := range occurrences {
        values = append(values, []int{key, value})
    }
    ret := make([]int, k)
    qsort(values, 0, len(values) - 1, ret, 0, k)
    return ret
}

func qsort(values [][]int, start, end int, ret []int, retIndex, k int) {
    rand.Seed(time.Now().UnixNano())
    picked := rand.Int() % (end - start + 1) + start;
    values[picked], values[start] = values[start], values[picked]

    pivot := values[start][1]
    index := start

    for i := start + 1; i <= end; i++ {
        if values[i][1] >= pivot {
            values[index + 1], values[i] = values[i], values[index + 1]
            index++
        }
    }
    values[start], values[index] = values[index], values[start]
    if k <= index - start {
        qsort(values, start, index - 1, ret, retIndex, k)
    } else {
        for i := start; i <= index; i++ {
            ret[retIndex] = values[i][0]
            retIndex++
        }
        if k > index - start + 1 {
            qsort(values, index + 1, end, ret, retIndex, k - (index - start + 1))
        }
    }
}