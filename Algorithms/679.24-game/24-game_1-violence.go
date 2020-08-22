func judgePoint24(nums []int) bool {
	//解法1，暴力法
	return judgePoint24_3(float64(nums[0])+float64(nums[1]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[0])-float64(nums[1]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[0])*float64(nums[1]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[0])/float64(nums[1]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[1])-float64(nums[0]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[1])/float64(nums[0]), float64(nums[2]), float64(nums[3])) ||

		judgePoint24_3(float64(nums[0])+float64(nums[2]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[0])-float64(nums[2]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[0])*float64(nums[2]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[0])/float64(nums[2]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[2])-float64(nums[0]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[2])/float64(nums[0]), float64(nums[1]), float64(nums[3])) ||

		judgePoint24_3(float64(nums[0])+float64(nums[3]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[0])-float64(nums[3]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[0])*float64(nums[3]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[0])/float64(nums[3]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[3])-float64(nums[0]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[3])/float64(nums[0]), float64(nums[2]), float64(nums[1])) ||

		judgePoint24_3(float64(nums[2])+float64(nums[3]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[2])-float64(nums[3]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[2])*float64(nums[3]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[2])/float64(nums[3]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[3])-float64(nums[2]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_3(float64(nums[3])/float64(nums[2]), float64(nums[0]), float64(nums[1])) ||

		judgePoint24_3(float64(nums[1])+float64(nums[2]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[1])-float64(nums[2]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[1])*float64(nums[2]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[1])/float64(nums[2]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[2])-float64(nums[1]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_3(float64(nums[2])/float64(nums[1]), float64(nums[0]), float64(nums[3])) ||

		judgePoint24_3(float64(nums[1])+float64(nums[3]), float64(nums[2]), float64(nums[0])) ||
		judgePoint24_3(float64(nums[1])-float64(nums[3]), float64(nums[2]), float64(nums[0])) ||
		judgePoint24_3(float64(nums[1])*float64(nums[3]), float64(nums[2]), float64(nums[0])) ||
		judgePoint24_3(float64(nums[1])/float64(nums[3]), float64(nums[2]), float64(nums[0])) ||
		judgePoint24_3(float64(nums[3])-float64(nums[1]), float64(nums[2]), float64(nums[0])) ||
		judgePoint24_3(float64(nums[3])/float64(nums[1]), float64(nums[2]), float64(nums[0]))
}

func judgePoint24_3(a, b, c float64) bool {
	return judgePoint24_2(a+b, c) ||
		judgePoint24_2(a-b, c) ||
		judgePoint24_2(a*b, c) ||
		judgePoint24_2(a/b, c) ||
		judgePoint24_2(b-a, c) ||
		judgePoint24_2(b/a, c) ||

		judgePoint24_2(a+c, b) ||
		judgePoint24_2(a-c, b) ||
		judgePoint24_2(a*c, b) ||
		judgePoint24_2(a/c, b) ||
		judgePoint24_2(c-a, b) ||
		judgePoint24_2(c/a, b) ||

		judgePoint24_2(c+b, a) ||
		judgePoint24_2(c-b, a) ||
		judgePoint24_2(c*b, a) ||
		judgePoint24_2(c/b, a) ||
		judgePoint24_2(b-c, a) ||
		judgePoint24_2(b/c, a)
}

func judgePoint24_2(a, b float64) bool {
	return (a+b < 24+1e-6 && a+b > 24-1e-6) ||
		(a*b < 24+1e-6 && a*b > 24-1e-6) ||
		(a-b < 24+1e-6 && a-b > 24-1e-6) ||
		(b-a < 24+1e-6 && b-a > 24-1e-6) ||
		(a/b < 24+1e-6 && a/b > 24-1e-6) ||
		(b/a < 24+1e-6 && b/a > 24-1e-6)
}

