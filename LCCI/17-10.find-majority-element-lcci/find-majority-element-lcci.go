func majorityElement(nums []int) int {
    //记录元素的个数
    var nummap = make(map[int]int)
    for _,num := range nums {
        nummap[num]++
    }
    //从map中找出主要元素
    for k,v := range nummap {
        //这里需要考虑 3/2 == 1这样的情况
        if float32(v) > float32(len(nums)/2.0) {
            return k
        }
    }

    return -1
}