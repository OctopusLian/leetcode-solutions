func missingNumber(nums []int) int {
    //等差数列解法
    //等差数列前 n 项和:Sn = n * (a1+an) / 2
    n := len(nums)  //数组元素总数
    sum := (n+1) * n/2
    for _, num := range(nums){
        sum -=  num
    }
    return sum  //剩下的值就是nums数组缺失的数字，返回即可
}
