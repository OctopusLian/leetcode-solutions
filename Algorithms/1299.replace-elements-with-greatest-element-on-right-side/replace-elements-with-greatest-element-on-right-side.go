func replaceElements(arr []int) []int {
    n := len(arr)
    max := arr[n-1]
    for i := n-2; i >= 0; i-- {
        cur := arr[i]
        arr[i] = max
        if max < cur {
            max = cur
        }
    }
    
    arr[n-1] = -1
    
    return arr
}