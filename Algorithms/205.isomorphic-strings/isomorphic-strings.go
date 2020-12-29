func isIsomorphic(s string, t string) bool {
    s2t := map[byte]byte{}
    t2s := map[byte]byte{}
    for i := range s {
        x, y := s[i], t[i]
        if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
            return false
        }
        s2t[x] = y
        t2s[y] = x
    }
    return true
}