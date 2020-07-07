func pow3(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res = (res * 3) % 1000000007
	}
	return res
}

func cuttingRope(n int) int {
	//解法1，贪心算法
    if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}

	if n%3 == 0 {
		return pow3(n / 3)
	} else if n%3 == 1 {
		return pow3((n-3)/3) * 4% 1000000007
	} else if n%3 == 2 {
		return pow3((n-2)/3) * 2% 1000000007
	}
	return 0
}