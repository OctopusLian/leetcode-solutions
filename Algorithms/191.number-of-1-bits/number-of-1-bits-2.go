func hammingWeight(num uint32) int {
    bits := 0
    mask := 1
    for i:=0;i<32;i++ {
        if ((int(num) & mask) != 0) {
            bits++
        }
        mask <<= 1
    }
    return bits
}
