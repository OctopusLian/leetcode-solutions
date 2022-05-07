/**
 * @param A: An integer
 * @return: a float number
 */
 func maxOfArray (A []float32) float32 {
    max := A[0]
    for i:=1;i<len(A);i++{
        if A[i] > max {
            max = A[i]
        }
    }
    return max
}