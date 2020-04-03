import 	"strconv"

const (
	Init    = iota //初始状态
	Space          //空格
	Number         //数字
	AriChar        //算数符号
	Zero
)

func myAtoi(str string) int {
	//状态机
	state := Init
	var beginIndex = -1
	var endIndex = -1
	var i = -1
	var c int32
	positive := true
loop:
	for i, c = range str {
		switch state {
		case Init:
			if isSpace(c) {
				state = Space;
				continue
			} else if (c == '0') {
				state = Zero
				beginIndex = i
				continue
			} else if isNumber(c) {
				state = Number;
				beginIndex = i
				continue
			} else if isAriChar(c) {
				positive = c == '+'
				state = AriChar;
				continue
			}
			return 0
		case Space:
			if isSpace(c) {
				continue
			} else if (c == '0') {
				state = Zero
				beginIndex = i
				continue
			} else if isNumber(c) {
				state = Number;
				beginIndex = i
				continue
			} else if isAriChar(c) {
				positive = c == '+'
				state = AriChar;
				beginIndex = i
				continue
			}
			return 0
		case AriChar:
			if (c == '0') {
				state = Zero
				beginIndex = i
				continue
			} else if isNumber(c) {
				beginIndex = i
				state = Number
				continue
			}
			return 0
		case Number:
			if isNumber(c) {
				continue
			}
			endIndex = i
			break loop
		case Zero:
			if c == '0' {
				beginIndex = i
				continue
			} else if (isNumber(c)) {
				state = Number
				beginIndex = i
				continue
			}
			return 0
		}
	}

	if beginIndex == -1 || i == -1 {
		return 0
	}
	if endIndex == -1 {
		endIndex = i + 1
	}

	str = str[beginIndex:endIndex]
	//str = strings.TrimLeft(str, "0")

	if !positive && strIntNumberCompare(str, "2147483648") == 1 {
		return -2147483648
	} else if positive && strIntNumberCompare(str, "2147483647") == 1 {
		return 2147483647
	}

	result, _ := strconv.Atoi(str)
	if positive {
		return result
	}

	return -result
}

func isSpace(c int32) bool {
	return c == ' '
}

func isNumber(c int32) bool {
	return c >= '0' && c <= '9'
}

func isAriChar(c int32) bool {
	return c == '-' || c == '+'
}

func strIntNumberCompare(a, b string) int {
	if len(a) > len(b) {
		return 1
	}
	if len(a) < len(b) {
		return -1
	}
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	case a == b:
		return 0
	}

	return -1
}
