func longestMountain(A []int) int {
    max, now := 0, 0
	asc, desc := false, false
	for i := 1; i < len(A); i++ {
		if asc {
			if A[i] > A[i-1] {
				now++
			} else if A[i] < A[i-1] {
				asc = false
				desc = true
				now++
			} else {
				now = 1
				asc = false
			}
		} else if desc {
			if A[i] < A[i-1] {
				now++
			} else {
				asc = true
				desc = false
				now = 1
				if A[i] > A[i-1] {
					now = 2
				}
			}
		} else if A[i] > A[i-1] {
			asc = true
			now = 2
		}

		if desc {
			if now > max {
				max = now
			}
		}
	}
	return max
}