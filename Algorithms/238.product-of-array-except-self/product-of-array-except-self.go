func productExceptSelf(nums []int) []int {
    // The length of the input array
    length := len(nums)
    // The left and right arrays as described in the algorithm
    var L = make([]int,length)
    var R = make([]int,length)
    // Final answer array to be returned
    var answer = make([]int,length)
    // L[i] contains the product of all the elements to the left
    // Note: for the element at index '0', there are no elements to the left,
    // so L[0] would be 1
    L[0] = 1
    for i := 1;i < length;i++ {
        // L[i - 1] already contains the product of elements to the left of 'i - 1'
        // Simply multiplying it with nums[i - 1] would give the product of all
        // elements to the left of index 'i'
        L[i] = nums[i-1] * L[i-1]
    }
    // R[i] contains the product of all the elements to the right
    // Note: for the element at index 'length - 1', there are no elements to the right,
    // so the R[length - 1] would be 1
    R[length - 1] = 1
    for i := length - 2;i >= 0;i-- {
        // R[i + 1] already contains the product of elements to the right of 'i + 1'
        // Simply multiplying it with nums[i + 1] would give the product of all
        // elements to the right of index 'i'
        R[i] = nums[i + 1] * R[i + 1]
    }
    // Constructing the answer array
    for i := 0;i < length;i++ {
        // For the first element, R[i] would be product except self
        // For the last element of the array, product except self would be L[i]
        // Else, multiple product of all elements to the left and to the right
        answer[i] = L[i] * R[i]
    }

    return answer
}
