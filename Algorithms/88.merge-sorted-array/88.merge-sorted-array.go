/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */
func merge(nums1 []int, m int, nums2 []int, n int)  {
	nums := make([]int,n+m)  //创建一个混合数组，长度是nums1和nums2的和
	//三个变量i，j，k，分别指向nums1，nums2，和混合数组的末尾
	for i,j,k := 0,0,0;i < m || j < n;k++ {
		if i < m && j < n {
			if nums1[i] < nums2[j] {
				nums[k] = nums1[i]
				i++
			} else {
				nums[k] = nums2[j]
				j++
			}
		} else if i < m {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
	}
	copy(nums1,nums)
}

