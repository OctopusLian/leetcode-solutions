func smallestK(arr []int, k int) []int {
    if k == 0 || len(arr) == 0{
        return []int{}
    }

    if len(arr) == 1 {
        return arr
    }

    res := []int{}
    left := []int{}
    right := []int{}
    tmp := arr[0]
    for i:=1;i<len(arr);i++ {
        if arr[i] < tmp {
            left = append(left, arr[i])
        } else {
            right = append(right, arr[i])
        }
    }

    if len(left) >= k {
        res = append(res, smallestK(left, k)...)
    } else if len(left) == k - 1{
        res = append(res, smallestK(left, k)...)
        res = append(res, tmp)
    } else {
        res = append(res, smallestK(left, len(left))...)
        res = append(res, tmp)
        res = append(res, smallestK(right, k - len(left) - 1)...)
    }

    return res
}