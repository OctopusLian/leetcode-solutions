func majorityElement(nums []int) int {
    res,cnt := 0,0  //初始化两个变量，res为返回的最终结果，cnt为出现次数
    for _,num := range nums {
        if cnt == 0 {
            //赋值数组中的第一个元素，并将cnt做+1处理
            res = num
            cnt += 1
        } else {
            if num == res {  //如果现在的数和前一个数相等
                cnt += 1 //出现次数+1
            } else {
                cnt -= 1  //否则出现次数-1，当减为0时res赋值给下一个元素，继续判断此时res的值是否是出现次数超过一半的数字
            }
        }
    }
    return res
}
