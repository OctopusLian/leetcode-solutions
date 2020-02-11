func rob(nums []int) int {
    preMax := 0
    currMax := 0
    for _,num := range nums {
        temp := currMax
        currMax = GetMax(preMax + num,currMax)
        preMax = temp
    }

    return currMax
}

func GetMax(x int,y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}
