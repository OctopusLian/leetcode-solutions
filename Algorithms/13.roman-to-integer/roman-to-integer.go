func romanToInt(s string) int {
    var charToIntMap = make(map[byte]int, 8)
	charToIntMap['I'] = 1
	charToIntMap['V'] = 5
	charToIntMap['X'] = 10
	charToIntMap['L'] = 50
	charToIntMap['C'] = 100
	charToIntMap['D'] = 500
	charToIntMap['M'] = 1000
    sum := 0
	for i:= 0; i < len(s); i++ {
        switch s[i] {
            case 'I':
            if i + 1 < len(s) && (s[i+1] == 'V' || s[i+1] == 'X') {
                sum -= charToIntMap[s[i]]
            } else {
                sum += charToIntMap[s[i]]
            }
            break

            case 'X':
			if i + 1 < len(s) && (s[i+1] == 'L' || s[i+1] == 'C') {
				sum -= charToIntMap[s[i]]
			} else {
				sum += charToIntMap[s[i]]
			}
			break

            case 'C':
			if i + 1 < len(s) && (s[i+1] == 'D' || s[i+1] == 'M') {
				sum -= charToIntMap[s[i]]
			} else {
				sum += charToIntMap[s[i]]
			}
			break

            default:
			sum += charToIntMap[s[i]]
        }
    }

    return sum
}
