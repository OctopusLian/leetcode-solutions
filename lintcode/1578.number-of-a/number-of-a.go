/**
 * @param s: the given string
 * @return: the number of A
 */
 func countA (s string) int {
    res := 0
    for _,v := range s {
        if v == 'A' {
            res++
        }
    }
    
    return res
}
