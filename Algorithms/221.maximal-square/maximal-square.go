func maximalSquare(matrix [][]byte) int {
    //1，暴力法
    maxSide := 0
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return maxSide
    }
    rows, columns := len(matrix), len(matrix[0])
    for i := 0; i < rows; i++ {
        for j := 0; j < columns; j++ {
            if matrix[i][j] == '1' {
                maxSide = max(maxSide, 1)
                curMaxSide := min(rows - i, columns - j)
                for k := 1; k < curMaxSide; k++ {
                    flag := true
                    if matrix[i+k][j+k] == '0' {
                        break
                    }
                    for m := 0; m < k; m++ {
                        if matrix[i+k][j+m] == '0' || matrix[i+m][j+k] == '0' {
                            flag = false
                            break
                        }
                    }
                    if flag {
                        maxSide = max(maxSide, k + 1)
                    } else {
                        break
                    }
                }
            }
        }
    }
    return maxSide * maxSide
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