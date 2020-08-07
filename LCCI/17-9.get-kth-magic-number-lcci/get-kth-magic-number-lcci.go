func getKthMagicNumber(k int) int {
    nums:=make([]int,k)
	nums[0]=1
	p3,p5,p7:=0,0,0
	for i:=1;i<k;i++{ 
		nums[i]=int(math.Min(float64(nums[p3]*3),math.Min(float64(nums[p5]*5),float64(nums[p7]*7))))
		if nums[i]==nums[p3]*3{
			p3++
		}
		if nums[i]==nums[p5]*5{
			p5++
		}
		if nums[i]==nums[p7]*7{
			p7++
		}
	}
	return nums[k-1]
}