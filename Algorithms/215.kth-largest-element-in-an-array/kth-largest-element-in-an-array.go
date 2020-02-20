func findKthLargest(nums []int, k int) int {
    //先截取最前面k个元素，组成一个新的list, O(k)
    queue := nums[0:k]
    //对该list进行升序排列 O(klog(k))
    sort.Ints(queue)

    //从nums的第k+1个元素，向后遍历, O(k*klog(k))
    for i := k; i < len(nums); i++ {
        //若该元素比list[0]小，则忽略（因为我们要取第k大的，所以，如果它比已知的k个元素中最小的还小，那么就可以忽略了）
        if nums[i] <= queue[0] {
            continue
        }else{
            //若该元素比已知的前k个元素最小的一个大，则替换掉前k个元素中最小的那个值，并重新排序前k个数
            queue[0] = nums[i]
            sort.Ints(queue)
        }
    }
    return queue[0]
}
