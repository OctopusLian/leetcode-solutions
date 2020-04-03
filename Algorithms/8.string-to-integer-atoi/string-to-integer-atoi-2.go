const (
    int32Max = 1<<31 - 1
    int32Min = -1 << 31
)

func myAtoi(str string) int {
    //单指针遍历串
    n := len(str)
    var i, j int
    neg := false
    for i = 0; i < n; i++ {
        if str[i] >= '0' && str[i] <= '9' {
            break
        } else if str[i] == '+' {
            i++
            break
        } else if str[i] == '-' {
            neg = true
            i++
            break
        } else if str[i] != ' ' {
            return 0
        }
    }
    for j = i; j < n; j++ {
        if str[j] < '0' || str[j] > '9' {
            break
        }
    }
    ret := 0
    for k := i; k < j; k++ {
        cur := int(str[k] - '0')
        if !neg {
            ret = ret*10 + cur
            if ret > int32Max {
                return int32Max
            }
        } else {
            ret = ret*10 - cur
            if ret < int32Min {
                return int32Min
            }
        }
    }
    return ret
}
