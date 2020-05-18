var s string

func permutation(S string) []string {
	//解法1，dfs
	s = S

	res := []string{}
	for i := range s {
		res = append(res, dfs("", i, make(map[int]bool))...)
	}
	return res
}

func dfs(temp string, now int, flag map[int]bool) []string {
	if now < 0 || now >= len(s) {
		return nil
	}
	if flag[now] {
		return nil
	}
	flag[now] = true
	temp += string(s[now])
	if len(temp) == len(s) {
		return []string{temp}
	}
	res := []string{}
	for i := range s {
		if !flag[i] {
			res = append(res, dfs(temp, i, flag)...)
			flag[i] = false
		}
	}
	return res
}