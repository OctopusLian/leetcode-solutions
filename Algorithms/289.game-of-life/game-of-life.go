func gameOfLife(board [][]int)  {
    //方法1,复制原数组并模拟
    var neighbors [3]int = [3]int{0,1,-1}
    rows := len(board)  //行数
    cols := len(board[0])  //列数

    // 创建复制数组 copyBoard
    copyBoard := make([][]int,len(board))

    //从原数组复制到一份到copyBoard
    for row := 0;row < rows;row++ {
        for col := 0;col < cols;col++ {
            copyBoard[row][col] = board[row][col]
        }
    }

    //遍历面板每一个格子里的细胞
    for row := 0;row < rows;row++ {
        for col := 0;col < cols;col++ {
            // 对于每一个细胞统计其八个相邻位置里的活细胞数量
            liveNeighbors := 0
        
            for i := 0;i < 3;i++ {
                for j := 0;j < 3;j++ {
                    if (!(neighbors[i] == 0 && neighbors[j] == 0)) {
                        r := row + neighbors[i]
                        c := col + neighbors[j]
                        // 查看相邻的细胞是否是活细胞
                            if ((r < rows && r >= 0) && (c < cols && c >= 0) && (copyBoard[r][c] == 1)) {
                                liveNeighbors += 1;
                            }
                    }
                    // 规则 1 或规则 3      
                    if ((copyBoard[row][col] == 1) && (liveNeighbors < 2 || liveNeighbors > 3)) {
                        board[row][col] = 0
                    }
                    // 规则 4
                    if (copyBoard[row][col] == 0 && liveNeighbors == 3) {
                        board[row][col] = 1
                    }
                }
            }
        }
    }
}
