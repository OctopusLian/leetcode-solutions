func canThreePartsEqualSum(A []int) bool {
    //双指针法
    length := len(A)
    left,right := 0,length-1
    leftSum,rightSum,sum := A[0],A[length-1],0
    for i:=0;i<len(A);i++ {
        sum += A[i]
    }
    for (left+1)<right {
        if (leftSum == rightSum) && (rightSum * 3 == sum) {
            return true
        }
        if (leftSum != sum/3) {
            left=left+1
            leftSum += A[left]
        }
        if (rightSum != sum/3) {
            right=right-1
            rightSum += A[right]
        }
    }
    return false
}
