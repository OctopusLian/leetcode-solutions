func intersect(nums1 []int, nums2 []int) []int {
    //考虑特殊情况
    if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
    //用两个字典，分别记录两个数组里元素出现的个数；键为元素，值为其出现的个数
	m1 := make(map[int]int, len(nums1))
	for _, v := range nums1 {
		m1[v] ++
	}
	m2 := make(map[int]int, len(nums2))
	for _, v := range nums2 {
		m2[v] ++
	}
    //合并结果
    var result []int
    for k1,v1 := range m1 {
        if m2[k1] == 0 {
            continue
        }
        for i := 0; i < min(v1, m2[k1]); i++ {
			result = append(result, k1)
		}
    }
    return result
}

func min(x int,y int) int {
    if x < y {
        return x
    } else {
        return y
    }
}
