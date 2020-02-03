func titleToNumber(s string) int {
    //将字符串变为字符数组
    var data []byte = []byte(s)
	ans:=0
	for i:=0;i<len(s);i++ {
		sum:=int(data[i])-64  //ASCII码
		ans = ans*26 + sum
	}
	return ans
}
