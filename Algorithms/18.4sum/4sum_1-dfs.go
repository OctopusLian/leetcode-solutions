func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)//先排序，方便之后的去重
	fourSum_helper(nums,target,[]int{},&res,0)
	return res
}

func fourSum_helper(nums []int,target int,cur []int, res *[][]int,index int){
	if len(cur)==4{
		if cur[0]+cur[1]+cur[2]+cur[3]==target{//这里只要改成n就可以解决n数相加问题
			temp:=make([]int,len(cur))
			copy(temp,cur)
			*res=append(*res,temp)
		}
		return
	}
	r:=math.MaxInt64//临时变量，解决同一深度的重复问题
	for i := index; i <len(nums); i++ {
		if r!=nums[i]{//剪枝操作
			r=nums[i]
			cur=append(cur,nums[i])
			fourSum_helper(nums,target,cur,res,i+1)
			cur=cur[:len(cur)-1]
		}
	}
}