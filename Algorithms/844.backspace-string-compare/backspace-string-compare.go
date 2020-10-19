func backspaceCompare(S string, T string) bool {
	return getFormatStr(S) == getFormatStr(T)
}

func getFormatStr(s string) string {
	l := len(s)
	right := 0
	sBytes := []byte(s)
	for i := 0; i < l; i++ {
		if sBytes[i] == "#"[0] {
			right--
			if right < 0 {
				right = 0
			}
		} else {
			sBytes[right] = sBytes[i]
			right++
		}
	}
	return string(sBytes[:right])
}