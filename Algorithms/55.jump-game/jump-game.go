func canJump(nums []int) bool {
	if len(nums) < 1 {
		return false
	}

	// 自己跳自己
	if len(nums) == 1 {
		return true
	}

	// 一开始就跳不动了
	if nums[0] == 0 {
		return false
	}

	zero := -1
	for i := len(nums) - 2; i >= 0; i-- {
		// 已经有0值了
		if zero > 0 {
			// 可以跳过0值
			if i+nums[i] > zero {
				// 当前0值可以忽略
				zero = -1
			}

			continue
		}

		if nums[i] == 0 {
			zero = i
			continue
		}

	}

	if zero < 0 {
		return true
	}

	return false
}
