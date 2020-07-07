func strToInt(str string) int {
	//解法1，正则表达式
    max := int(math.Pow(float64(2), float64(31))) - 1
    min := -int(math.Pow(float64(2), float64(31)))
    re := regexp.MustCompile(`^[+-]?\d+`)
    slice := re.FindAllString(strings.TrimSpace(str), -1)
    str = strings.Join(slice, "")
    res, _ := strconv.Atoi(str)
    if res < min {
        res = min
    } else if res > max {
        res = max
    }
    return res 
}