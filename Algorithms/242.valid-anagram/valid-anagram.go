func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    countmap := make(map[uint8]int)
    for i := 0; i < len(s); i++ {
		countmap[s[i]]++
	}
    for i := 0; i < len(t); i++ {
        countmap[t[i]]--
        if countmap[t[i]] < 0 {
            return false
        }
    }

    return true
}
