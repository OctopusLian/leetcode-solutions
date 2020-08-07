func isNumber(s string) bool {
	//解法1，正则表达式
    s = strings.TrimSpace(s)
    res,_ := regexp.MatchString("^(([\\+\\-]?[0-9]+(\\.[0-9]*)?)|([\\+\\-]?\\.?[0-9]+))(e[\\+\\-]?[0-9]+)?$", s)

    return res
}