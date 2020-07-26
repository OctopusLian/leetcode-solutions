var (
    dirs = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}
    rows, columns int
)

func longestIncreasingPath(matrix [][]int) int {
	//解法1，拓扑排序
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }
    rows, columns = len(matrix), len(matrix[0])
    outdegrees := make([][]int, rows)
    for i := 0; i < rows; i++ {
        outdegrees[i] = make([]int, columns)
    }
    for i := 0; i < rows; i++ {
        for j := 0; j < columns; j++ {
            for _, dir := range dirs {
                newRow, newColumn := i + dir[0], j + dir[1]
                if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] > matrix[i][j] {
                    outdegrees[i][j]++
                }
            }
        }
    }

    queue := [][]int{}
    for i := 0; i < rows; i++ {
        for j := 0; j < columns; j++ {
            if outdegrees[i][j] == 0 {
                queue = append(queue, []int{i, j})
            }
        }
    }
    ans := 0
    for len(queue) != 0 {
        ans++
        size := len(queue)
        for i := 0; i < size; i++ {
            cell := queue[0]
            queue = queue[1:]
            row, column := cell[0], cell[1]
            for _, dir := range dirs {
                newRow, newColumn := row + dir[0], column + dir[1]
                if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] < matrix[row][column] {
                    outdegrees[newRow][newColumn]--
                    if outdegrees[newRow][newColumn] == 0 {
                        queue = append(queue, []int{newRow, newColumn})
                    }
                }
            }
        }
    }
    return ans
}