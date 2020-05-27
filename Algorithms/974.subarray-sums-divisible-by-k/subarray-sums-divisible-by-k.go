func subarraysDivByK(A []int, K int) int {
    record := map[int]int{0: 1}
    sum, ans := 0, 0
    for _, elem := range A {
        sum += elem
        modulus := (sum % K + K) % K
        ans += record[modulus]
        record[modulus]++
    } 
    return ans
}