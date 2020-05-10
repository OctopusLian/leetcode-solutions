func findString(words []string, s string) int {
    for index,word := range words {
        if word == s {
            return index
        }
    }

    return -1
}