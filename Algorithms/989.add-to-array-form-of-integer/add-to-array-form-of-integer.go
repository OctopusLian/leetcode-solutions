func addToArrayForm(A []int, K int) []int {
    llen := len(A)
	A[llen-1] += K
	carry := 0
	result := []int{}
	for i := llen -1 ; i >= 0 ; i -- {
		carry = A[i] / 10
		A[i] = A[i] % 10
		if i > 0 {
			A[i-1] += carry
		}
	}
	if carry != 0 {
		for _,x := range strings.Split(strconv.Itoa(carry),"") {
			z,_ := strconv.ParseInt(x,10,64)
			result = append(result, int(z))
		}
	}
	result = append(result,A...)
	return result
}