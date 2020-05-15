func insertBits(N int, M int, i int, j int) int {
    return ((N >> (j + 1)) << (j + 1)) | ((N >> i << i) ^ N)|( M<<i)
}