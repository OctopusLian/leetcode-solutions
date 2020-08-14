func findClosedNumbers(num int) []int {
    hasBig, hasSmall := false, false
	big, small := num+1, num-1
	b1c :=  findBit1Count(num)
	if b1c == 0 {
		return []int{-1, -1}
	}
	for big < 2147483647 {
		if findBit1Count(big) == b1c {
			hasBig = true
			break
		}
		big++
	}
	for small > 0 {
		if findBit1Count(small) == b1c {
			hasSmall = true
			break
		}
		small--
	}
	if !hasBig {
		big = -1
	}
	if !hasSmall {
		small = -1
	}
	return []int{big, small}
}


func findBit1Count(num int) int  {
	count := 0
	for num != 0 {
		if num & 1 != 0 {
			count++
		}
		num >>= 1
	}
	return count
}