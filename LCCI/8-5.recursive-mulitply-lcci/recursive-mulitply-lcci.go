func multiply(A int, B int) int {
    if B==0{
        return 0
    }
    return multiply(A,B-1)+A
}