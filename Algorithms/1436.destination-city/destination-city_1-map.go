func destCity(paths [][]string) string {
    front := make(map[string]bool)
    end   := make(map[string]bool)
    for _, path := range paths {
        // 记录所有出现过在起点的元素
        front[path[0]] = true
        // 记录所有出现过在终点的元素
        end[path[1]]   = true
    }

    // 遍历所有出现过在终点的元素，如果某元素只出现过在终点，没有出现过在起点，则该元素就是最终结果
    for key, _ := range end {
        if _, ok := front[key]; !ok {
            return key
        }
    }
    return ""
}