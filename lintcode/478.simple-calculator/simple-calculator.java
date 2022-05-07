public class Calculator {
    /**
     * @param a:        An integer
     * @param operator: A character, +, -, *, /.
     * @param b:        An integer
     * @return: The result
     */
    public int calculate(int a, char operator, int b) {
        // write your code here
        switch (operator) {
        case '+':
            return a + b;
        case '-':
            return a - b;
        case '*':
            return a * b;
        case '/':
            return a / b;
        }
        return 0;
    }
}