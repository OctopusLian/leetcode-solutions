func findClosest(words []string, word1 string, word2 string) int {
    memo := make(map[string]int)  //建立word map
    results := 1 << 31 -1
    for idx, word := range words {
        if i, ok := memo[word]; ok {
            if results > idx - i {
                results = idx - i
            }
        }
        if word == word1 {
            memo[word2] = idx // 反着存，这样可以in
        }

        if word == word2 {
            memo[word1] = idx
        }
    }
    return results
}