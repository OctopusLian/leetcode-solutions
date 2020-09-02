// 全局变量,保存结果
var result [][]string

// 回溯核心
// board： 棋盘
// path：选择列表
func backtrack(board [][]bool, path [][]byte) {
	// 结束条件
	if len(path) == len(board) {
		t := make([]string, len(path))
		for k, bs := range path {
			t[k] = string(bs)
		}
		result = append(result, t)
	}

	for i := 0; i < len(board); i++ {
		// 不合法情况,剔除(剪枝)
		if !isValid(board, len(path), i) {
			continue
		}
		// 打印默认位置
		bs := printLine(len(board))
		// 放置皇后
		bs[i] = 'Q'
		// 放入棋盘
		board[len(path)][i] = true
		// 加入选择路径
		path = append(path, bs)
		// 进行下一次决策
		backtrack(board, path)
		// 撤销选择
		path = path[:len(path)-1]
		board[len(path)][i] = false
	}
}

// 是否能在 board[row][col] 位置放置皇后
// 皇后不可以上下左右对角线同时存在
func isValid(board [][]bool, row, col int) bool {
	// 检查行是否有皇后冲突
	for i := 0; i < row; i++ {
		if board[i][col] == true {
			return false
		}
	}
	// 检查列是否有皇后冲突
	for j := 0; j < len(board); j++ {
		if board[row][j] == true {
			return false
		}
	}
	// 检查对角线: "\"
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == true {
			return false
		}
	}
	// 检查对角线: "/"
	for i, j := row, col; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == true {
			return false
		}
	}

	return true
}

// 打印每行默认情况,都是'.'
func printLine(n int) []byte {
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = '.'
	}
	return bs
}

func solveNQueens(n int) [][]string {
	// 清空变量
	result = [][]string{}
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}
	backtrack(board, [][]byte{})
	return result
}