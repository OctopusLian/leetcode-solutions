/**
 * @param A: The array A.
 * @return: The array of the squares.
 */
 func SquareArray (A []int) []int {
    squareArraySort(A, 0, len(A)-1)
    return A
}

func squareArraySort (A []int, low int, high int) {
    if len(A) < 1 {
        return
    }
    if low > high {
        return
    }
    index := getIndex(A, low, high)
    //A[index] = A[index] * A[index]
    squareArraySort(A, low, index-1)
    squareArraySort(A, index+1, high)
}

func getIndex (A []int, low int, high int) int {
    temp := A[low]
    if temp < 0 {
        temp = -1 * temp
    }
    for low < high {
        for low < high {
            if A[high] < 0 {
                A[high] = -1 * A[high]
            }
            if A[high] >= temp {
                high--
                continue
            }
            break
        }
        A[low] = A[high]
        for low < high {
            if A[low] < 0 {
                A[low] = -1 * A[low]
            }
            if A[low] <= temp {
                low++
                continue
            }
            break
        }
        A[high] = A[low]
    }
    A[low] = temp * temp
    return low
}