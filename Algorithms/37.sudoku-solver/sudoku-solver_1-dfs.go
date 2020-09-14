func solveSudoku(board [][]byte) {
	//解法1，深度优先搜索
    row_map, col_map, box_map, blanks := getMaps(board)
    dfs(board, row_map, col_map, box_map, blanks)
}

func dfs(board [][]byte, row_map, col_map, box_map []map[byte]bool, blanks [][2]int) bool {
    // 如果没有空白格子，就说明解开了
    if len(blanks) == 0 {
        return true
    }
    // 获取要填的空白格子
    x, y := blanks[0][0], blanks[0][1]
    // 尝试填入1-9
    for i := 1; i <= 9; i++ {
        n := '0' + byte(i)
        // 如果能够填入这个数
        if canPut(n, x, y, row_map, col_map, box_map) {
            // 填进去
            board[x][y] = '0' + byte(i)
            // 更新映射
            updateMaps(n, x, y, row_map, col_map, box_map, true)
            // 填下一个格子，如果下一个格子完成了，直接返回结果
            if dfs(board, row_map, col_map, box_map, blanks[1:]) {
                return true
            } else {
                // 如果填这个数无法完成，就还原，尝试填入下一个数
                updateMaps(n, x, y, row_map, col_map, box_map, false)
            }
        }
    }
    // 如果九个数都没法填入，恢复原状并返回false
    board[x][y] = '.'
    return false
}

// 初始化行/列/九宫格的映射，并记录空格子
func getMaps(board [][]byte) (a, b, c []map[byte]bool, d [][2]int) {
    col_map := make([]map[byte]bool, 9, 9)  // 列
    row_map := make([]map[byte]bool, 9, 9)  // 行
    box_map := make([]map[byte]bool, 9, 9)  // 九宫格
    blanks := make([][2]int, 0, 81) // 空格子
    for i := 0; i < 9; i++ {
        col_map[i] = make(map[byte]bool)
        row_map[i] = make(map[byte]bool)
        box_map[i] = make(map[byte]bool)
    }
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] != '.' {
                row_map[i][board[i][j]] = true
                col_map[j][board[i][j]] = true
                box_map[3*(i/3)+j/3][board[i][j]] = true
            } else {
                // 将空白点的索引添加到空白数组
                blanks = append(blanks, [2]int{i, j})
            }
        }
    }
    return row_map, col_map, box_map, blanks
}

// 更新映射
func updateMaps(n byte, x, y int, row_map, col_map, box_map []map[byte]bool, val bool) {
    row_map[x][n] = val
    col_map[y][n] = val
    box_map[3*(x/3)+y/3][n] = val
}

// 判断格子 i, j 能否填入 n
func canPut(n byte, i, j int, row_map, col_map, box_map []map[byte]bool) bool {
    return !row_map[i][n] && !col_map[j][n] && !box_map[3*(i/3)+j/3][n]
}