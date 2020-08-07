func wiggleSort(nums []int)  {
    for i := 0; i < len(nums)-1; i++ {
		if (i%2 == 1 && nums[i] < nums[i+1]) || (i%2 == 0 && nums[i] > nums[i+1]) {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}
}