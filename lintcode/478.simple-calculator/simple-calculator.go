/**
 * @param a: An integer
 * @param operator: A character, +, -, *, /.
 * @param b: An integer
 * @return: The result
 */
 func calculate (a int, operator byte, b int) int {
    // write your code here
    switch operator {
            case '+': return a + b
            case '-': return a - b
            case '*': return a * b
            case '/': return a / b
    }
    return 0
}