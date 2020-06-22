func patternMatching(pattern string, value string) bool {
	//解法1，枚举法
    countA, countB := 0, 0
    for i := 0; i < len(pattern); i++ {
        if pattern[i] == 'a' {
            countA++
        } else {
            countB++
        }
    }
    if countA < countB {
        countA, countB = countB, countA
        tmp := ""
        for i := 0; i < len(pattern); i++ {
            if pattern[i] == 'a' {
                tmp += "b"
            } else {
                tmp += "a"
            }
        }
        pattern = tmp
    }
    if len(value) == 0 {
        return countB == 0
    }
    if len(pattern) == 0 {
        return false
    }

    for lenA := 0; countA * lenA <= len(value); lenA++ {
        rest := len(value) - countA * lenA
        if (countB == 0 && rest == 0) || (countB != 0 && rest % countB == 0) {
            var lenB int
            if countB == 0 {
                lenB = 0
            } else {
                lenB = rest / countB
            }
            pos, correct := 0, true
            var valueA, valueB string
            for i := 0; i < len(pattern); i++ {
                if pattern[i] == 'a' {
                    sub := value[pos:pos+lenA]
                    if len(valueA) == 0 {
                        valueA = sub
                    } else if valueA != sub {
                        correct = false
                        break
                    }
                    pos += lenA
                } else {
                    sub := value[pos:pos+lenB]
                    if len(valueB) == 0 {
                        valueB = sub
                    } else if valueB != sub {
                        correct = false
                        break
                    }
                    pos += lenB
                }
            }
            if correct && valueA != valueB {
                return true
            }
        } 
    }
    return false
}