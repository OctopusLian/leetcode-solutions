func reverseWords(s string) string {
	ss := strings.Fields(s)
	var results []string
	for _, str := range ss {
		var shift_str []byte
		for i := 0; i < len(str); i++ {
			shift_str = append(shift_str, str[len(str)-i-1])
		}
		results = append(results, string(shift_str))
	}

	re := ""
	for index, value := range results {
		if index == 0 {
			re = re + value
		} else {
			re = re + " " + value
		}
	}

	return re
}
