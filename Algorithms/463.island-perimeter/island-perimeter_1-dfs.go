type pair struct{ x, y int }
var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func islandPerimeter(grid [][]int) (ans int) {
    n, m := len(grid), len(grid[0])
    var dfs func(x, y int)
    dfs = func(x, y int) {
        if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
            ans++
            return
        }
        if grid[x][y] == 2 {
            return
        }
        grid[x][y] = 2
        for _, d := range dir4 {
            dfs(x+d.x, y+d.y)
        }
    }
    for i, row := range grid {
        for j, v := range row {
            if v == 1 {
                dfs(i, j)
            }
        }
    }
    return
}