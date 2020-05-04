func findMagicIndex(nums []int) int {
    return magicFast(nums,0,len(nums)-1)
}

func magicFast(array []int,start int,end int) int {
    if end < start {
        return -1
    }
    midIndex := (start + end) / 2
    midValue := array[midIndex]
    if midValue == midIndex {
        return midIndex
    }
    
    //搜索左半部分
    leftIndex := Min(midIndex - 1,midValue)
    left := magicFast(array,start,leftIndex)
    if left >= 0{
        return left
    }

    //搜索右半部分
    rightIndex := Max(midIndex + 1,midValue)
    right := magicFast(array,rightIndex,end)

    return right
}

func Min(x int,y int)int {
    if x > y {
        return y
    } else {
        return x
    }
}

func Max(x int,y int)int {
    if x > y {
        return x
    } else {
        return y
    }
}