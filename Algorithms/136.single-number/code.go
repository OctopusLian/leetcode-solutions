func singleNumber(nums []int) int {
    numsmap := make(map[int]int)
    var result int
    for _,v := range nums {
        if numsmap[v] != 0 {
            numsmap[v] = 2  //有重复，value为2
        } else {
            numsmap[v] = 1  //无重复，value为1
        }

    }
    for k,v := range numsmap {
        if v == 1 {
            result = k  //返回那个只出现一次的数字
        }
    }

    return result
}
