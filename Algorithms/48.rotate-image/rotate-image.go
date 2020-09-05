func rotate(matrix [][]int)  {
    size := len(matrix)
	subsize := size - 1
	for i:=0; i<size/2; i++ {
		for j:=i; j<subsize-i; j++ {
			matrix[i][j], matrix[j][subsize-i] = matrix[j][subsize-i], matrix[i][j]
			matrix[i][j], matrix[subsize-i][subsize-j] = matrix[subsize-i][subsize-j], matrix[i][j]
			matrix[i][j], matrix[subsize-j][i] = matrix[subsize-j][i], matrix[i][j]
		}
	}
}