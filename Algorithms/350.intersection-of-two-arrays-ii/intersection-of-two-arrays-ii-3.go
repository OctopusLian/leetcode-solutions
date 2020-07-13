func intersect(nums1 []int, nums2 []int) []int {
	//解法3，排序
	//首先对两个数组进行排序
    sort.Ints(nums1)
	sort.Ints(nums2)
	//然后使用两个指针遍历两个数组，初始时，两个指针分别指向两个数组的头部。
    length1, length2 := len(nums1), len(nums2)
    index1, index2 := 0, 0

	//每次比较两个指针指向的两个数组中的数字，如果两个数字不相等，则将指向较小数字的指针右移一位，
	//如果两个数字相等，将该数字添加到答案，并将两个指针都右移一位。当至少有一个指针超出数组范围时，遍历结束。
    intersection := []int{}
    for index1 < length1 && index2 < length2 {
        if nums1[index1] < nums2[index2] {
            index1++
        } else if nums1[index1] > nums2[index2] {
            index2++
        } else {
            intersection = append(intersection, nums1[index1])
            index1++
            index2++
        }
    }
    return intersection
}