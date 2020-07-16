var (
    UNCOLORED, RED, GREEN = 0, 1, 2
)

func isBipartite(graph [][]int) bool {
	//解法1，广度优先搜索
    n := len(graph)
    color := make([]int, n)
    for i := 0; i < n; i++ {
        if color[i] == UNCOLORED {
            queue := []int{}
            queue = append(queue, i)
            color[i] = RED
            for i := 0; i < len(queue); i++ {
                node := queue[i]
                cNei := RED
                if color[node] == RED {
                    cNei = GREEN
                }
                for _, neighbor := range graph[node] {
                    if color[neighbor] == UNCOLORED {
                        queue = append(queue, neighbor)
                        color[neighbor] = cNei
                    } else if color[neighbor] != cNei {
                        return false
                    } 
                }
            }
        }
    }
    return true
}