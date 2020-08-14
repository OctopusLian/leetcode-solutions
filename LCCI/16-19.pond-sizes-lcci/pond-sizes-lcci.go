type point struct {
	m, n int
}
  
func pondSizes(land [][]int) (szs []int) {
	M, N := len(land), len(land[0])
	dy := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for m := 0; m < M; m++ {
		for n := 0; n < N; n++ {
			if land[m][n] == 0 {
				land[m][n] = -1
				q := []point{}
				fptr := 0
				q = append(q, point{m, n})
				for fptr != len(q) {
				  i0, j0 := q[fptr].m, q[fptr].n
				  fptr++
				  for x := 0; x < 8; x++ {
					i, j := i0+dy[x][0], j0+dy[x][1]
					if i > -1 && i < M && j > -1 && j < N && land[i][j] == 0 {
						land[i][j] = -1
						q = append(q, point{i, j})
					}
				  }
				}
				szs = append(szs, len(q))
			}
		}
	}
	sort.Ints(szs)
	return
}