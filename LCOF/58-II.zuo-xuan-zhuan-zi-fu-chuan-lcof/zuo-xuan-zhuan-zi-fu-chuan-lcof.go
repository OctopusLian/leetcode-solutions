func reverseLeftWords(s string, n int) string {
    return s[n:len(s)] + s[:n]
}
