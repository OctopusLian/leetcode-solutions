func printKMoves(K int) []string {
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}       // 四个方向依次为右、下、左、上
	curDir := 0     // 当前方向
	colors := make(map[Point]int)        // 0 白色，1 黑色
	edges := make([]int, 4)     // 四个边界依次为最左侧、最上方、最右侧、最下方
	dirChar := []byte{'R', 'D', 'L', 'U'}       // 用于将更新结果数组
	curPos := Point{}

	for i := 0; i < K; i++ {
		if colors[curPos] == 0 {    // 顺时针
			curDir = (curDir + 1) % 4
		} else {    // 逆时针
			curDir = (curDir + 3) % 4
		}
		colors[curPos] ^= 1     // 变色
		curPos = Point{curPos.x + dirs[curDir][0], curPos.y + dirs[curDir][1]}   // 前进一步

		updateEdge(edges, curPos)      // 更新顶点
	}

	rows, cols := edges[3] - edges[1] + 1, edges[2] - edges[0] + 1
	res := make([]string, rows)

	for i := edges[1]; i <= edges[3]; i++ {
		row := make([]byte, cols)
		for j := edges[0]; j <= edges[2]; j++ {
			color := byte('_')
			if colors[Point{i,j}] == 1 {
				color = byte('X')
			}
			row[j-edges[0]] = color
			if i == curPos.x && j == curPos.y {     // 填充最终位置
				row[j-edges[0]] = dirChar[curDir]
			}
		}
		res[i-edges[1]] = string(row)
	}

	return res
}

func updateEdge(edges []int, pos Point) {
	edges[0] = min(edges[0], pos.y)
	edges[1] = min(edges[1], pos.x)
	edges[2] = max(edges[2], pos.y)
	edges[3] = max(edges[3], pos.x)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type Point struct {
	x, y int
}