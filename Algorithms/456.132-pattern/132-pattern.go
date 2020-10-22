func find132pattern(nums []int) bool {
    stack := make([]int, 0)
    min := -1 << 63

    for k := len(nums) - 1; k >= 0; k--{
        n := nums[k]

        if n < min{
            return true
        }

        for len(stack) > 0  && n > stack[len(stack) - 1]{
            i := len(stack) - 1
            min = Max(min, stack[i])
            stack = stack[:i]
        }
        
        stack = append(stack, n)
    }

    return false
}

func Max(a, b int) int{
    if a > b{
        return a
    }
    return b
}