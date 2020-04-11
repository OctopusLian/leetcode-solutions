func superEggDrop(K int, N int) int {
    N++
	for base := 0; ; base++ {
		maxN := 0
		for i := 0; i <= K && i <= base; i++ {
			maxN += comb(i, base)
		}
		if maxN >= N {
			return base
		}
	}
}

// 计算组合数C^m_n
func comb(m, n int) int {
	if m > n {
		m = n
	}
	if m > n/2 {
		m = n - m
	}
	ans := 1
	for i := n - m + 1; i <= n; i++ {
		ans *= i
	}
	for i := 2; i <= m; i++ {
		ans /= i
	}
	return ans
}

