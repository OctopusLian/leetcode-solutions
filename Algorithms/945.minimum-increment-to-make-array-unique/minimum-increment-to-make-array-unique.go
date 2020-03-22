func minIncrementForUnique(A []int) int {
    if (len(A) < 1) {
        return 0
    }
    //排序，得到一个趋势向上的台阶
    sort.Ints(A)
    var res int
    f := A[0] //一开始选择，把标准定在第一级，后续的第二级跟它补齐缺失的高度。后续的第i级和前一级比。
    for i := 1; i < len(A); i++ {
        if A[i] <= f { // 如果第i级不高于它的前一级，则补平，同时还要高出前一级1单位
            res = res + f - A[i] + 1 
            f++  //补齐当前高度后，后一级以新的高度作为对齐标准，实际上此处应为 f = A[i] + f - A[i] + 1
        }else {
            f = A[i] //如果第i级高于它的前一级，不用补平，把f更新为当前高度，一遍后一级对齐
        }
    }
    return res
}
