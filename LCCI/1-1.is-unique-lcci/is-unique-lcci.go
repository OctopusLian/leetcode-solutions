func isUnique(astr string) bool {
    if len(astr) > 128 {
        return false
    }

    var char_set = make(map[byte]bool)
    for i:=0;i<len(astr);i++ {
        if (char_set[astr[i]]) {
            return false
        }
        char_set[astr[i]] = true
    }
    return true
}
