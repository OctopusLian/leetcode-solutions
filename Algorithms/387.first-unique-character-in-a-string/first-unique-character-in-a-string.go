func firstUniqChar(s string) int {
    // build hash map : character and how often it appears
    cmap := make(map[byte]int)
    for i:=0;i<len(s);i++ {
        c := s[i]
        cmap[c] = cmap[c] + 1
    }

    // find the index
    for i:=0;i<len(s);i++ {
        if cmap[s[i]] == 1 {
            return i
        }
    }

    return -1
}
