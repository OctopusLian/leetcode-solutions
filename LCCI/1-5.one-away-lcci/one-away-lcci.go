func oneEditAway(first string, second string) bool {
    if math.Abs(float64(len(first))-float64(len(second))) > 1 { //长度相差超过1就返回
		return false
	}
	length1, length2 := len(first), len(second)
	differ := 0 //记录两个字符串比较有多少次的不同
	i, j := 0, 0
	for i < length1 && j < length2 {
		if first[i] != second[j] {
			differ++
			if length1 > length2 { //出现不同的情况，就把长的那个下标++，正常来说，下一个会相等，并且之后的字符也会相等
				i++
				continue
			} else if length1 < length2 {
				j++
				continue
			}
		}
		i++
		j++
	}

	return differ <= 1
}