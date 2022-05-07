/*
 * @Description:
 * @Author: neozhang
 * @Date: 2021-11-14 20:44:52
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-11 21:52:57
 */
/**
 * @param A: An integer array
 * @return: An integer
 */
func singleNumber(A []int) int {
	Amap := make(map[int]int)
	for _, a := range A {
		Amap[a]++
	}
	for k, v := range Amap {
		if v == 1 {
			return k
		}
	}

	return -1
}