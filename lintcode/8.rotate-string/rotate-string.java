public class Solution {
    /**
     * @param str:    An array of char
     * @param offset: An integer
     * @return: nothing
     */
    public void rotateString(char[] str, int offset) {
        // write your code here
        if (str.length == 0 || offset == 0) {
            return;
        }
        offset = offset % str.length;// 实际偏移量
        for (int i = 0; i < offset; i++) {
            char temp = str[str.length - 1];
            for (int j = str.length - 1; j > 0; j--) {
                str[j] = str[j - 1];
            }
            str[0] = temp;
        }
    }
}