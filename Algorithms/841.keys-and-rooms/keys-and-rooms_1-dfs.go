var (
    num int
    vis []bool
)

func canVisitAllRooms(rooms [][]int) bool {
	//解法1，深度优先搜索
    n := len(rooms)
    num = 0
    vis = make([]bool, n)
    dfs(rooms, 0)
    return num == n
}

func dfs(rooms [][]int, x int) {
    vis[x] = true
    num++
    for _, it := range rooms[x] {
        if !vis[it] {
            dfs(rooms, it)
        }
    }
}