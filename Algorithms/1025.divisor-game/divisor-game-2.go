func divisorGame(N int) bool {
	//解法2，递推
    f := make([]bool, N + 5)
    f[1], f[2] = false, true
    for i := 3; i <= N; i++ {
        for j := 1; j < i; j++ {
            if i % j == 0 && !f[i - j] {
                f[i] = true
                break
            }
        }
    }
    return f[N]
}