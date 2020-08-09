var result []string

func restoreIpAddresses(s string) []string {
	//解法1，回溯算法
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
    result = []string{}
	track := make([]string, 0)
	backtrack(s, track, 1)
	return result
}

func backtrack(s string, track []string, key int) {
	// 结束条件
	// key为段数，一共4段，选完4段，同时字符串选完时结束
	if key == 5 {
		if len(s) == 0 {
			str := strings.Join(track, ".")
			result = append(result, str)
		}
		return
	}
	// 选择列表
	// 每一段最大选择3位
	for j := 1; j <= 3; j++ {
		if j <= len(s) {
			// 选1-3位数字
			v, err := strconv.Atoi(s[:j])
			if err == nil && v <= 255 {
				// 做选择
				//fmt.Printf("第 %d 段，选择 %d位 s: %s\n", key, j, s[:j])
				track = append(track, s[:j])
				str := s[j:]
				// 下一段选择
				backtrack(str, track, key+1)
				// 撤销选择
				track = track[:len(track)-1]
			}
            // 每一段只能为0，不能为01
            if v == 0 {
				break
			}
		}
	}
}