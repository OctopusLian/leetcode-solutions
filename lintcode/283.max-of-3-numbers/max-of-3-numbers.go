/**
 * @param num1: An integer
 * @param num2: An integer
 * @param num3: An integer
 * @return: an interger
 */
func maxOfThreeNumbers (num1 int, num2 int, num3 int) int {
    res := 0
    if num1 > num2 {
        res = num1
    } else {
        res = num2
    }
    
    if res > num3 {
        return res
    } else {
        return num3
    }
}
