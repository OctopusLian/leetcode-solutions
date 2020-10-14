func commonChars(A []string) []string {
    firstWordMap := make(map[rune]int)
    for _, char := range A[0] {
        firstWordMap[char]++
    }
    result := make([]string, 0)
    for i := 1; i < len(A); i++ {
        wordMap := make(map[rune]int)
        for _, char := range A[i] {
            wordMap[char]++
        }
        for char, firstWordMapCharCount := range firstWordMap {
            count, _ := wordMap[char]
            if firstWordMapCharCount > count {
                firstWordMap[char] = count
            }
        }
    }
    for char, count := range firstWordMap {
        for i := 0; i < count; i++ {
            result = append(result, string(char))
        }
    }
    return result
}