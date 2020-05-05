func masterMind(solution string, guess string) []int {
    res:=make([]int,2)
	soluMap:=make(map[byte]int)
	for i:=0;i< len(solution);i++{
		if solution[i]==guess[i]{
			res[0]++
		}
		soluMap[solution[i]]++
	}
	b:=0
	for i:=0;i< len(guess);i++{
		if _,ok:=soluMap[guess[i]];ok&&soluMap[guess[i]]>0{
			b++
			soluMap[guess[i]]--
		}
	}
	res[1]=b-res[0]
	return res
}