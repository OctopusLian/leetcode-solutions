func minArray(numbers []int) int {
    //双指针，二分法
    left,right := 0,len(numbers)-1
    for ;left < right; {
        mid := (left + right) / 2
        if (numbers[mid] > numbers[right]) {
            left = mid + 1
        } else if (numbers[mid] < numbers[right]) {
            right = mid
        } else {
            right = right - 1
        }
    }
    return numbers[left]
}
