func canVisitAllRooms(rooms [][]int) bool {
	//解法2，广度优先搜索
    n := len(rooms)
    num := 0
    vis := make([]bool, n)
    queue := []int{}
    vis[0] = true
    queue = append(queue, 0)
    for i := 0; i < len(queue); i++ {
        x := queue[i]
        num++
        for _, it := range rooms[x] {
            if !vis[it] {
                vis[it] = true
                queue = append(queue, it)
            }
        }
    }
    return num == n
}