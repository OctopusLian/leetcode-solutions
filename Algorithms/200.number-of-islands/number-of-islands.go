func numIslands(grid [][]byte) int {
    result := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == '1' {            //遍历岛屿数量
                BFS(grid, i, j)
                result ++
            }
        }
    }
    return result
}

//建立xy坐标轴的方式遍历一个点的四个方向，并将遇到的岛屿转化成水域。递归退出的临界条件是遇到边界就退出，或者传入岛屿皆为0就直接退出。
var fx = [4]int{1, 0, -1, 0}     //x轴
var fy = [4]int{0, -1, 0, 1}     //y轴

func BFS(grid [][]byte, x int, y int){
    if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '0'{ //递归退出临界条件
        return
    }
    grid[x][y] = '0'                   //将传入的岛屿转化成为水域
    for i := 0; i < 4; i++ {           //通过坐标轴遍历一个点的四个方向
        BFS(grid,x + fx[i], y + fy[i]) 
    }
}
