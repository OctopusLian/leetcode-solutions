func surfaceArea(grid [][]int) int {
    area := 0
	for i, row := range grid {
		for j, level := range row {
			if level > 0 {
				area += level*4 + 2
				if i > 0 {
					area -= findMin(grid[i-1][j], level) * 2
				}
				if j > 0 {
					area -= findMin(grid[i][j-1], level) * 2
				}
			}
		}
	}
	return area
}

func findMin(i, j int) int {
	if i > j {
		return j
	}
	return i
}
