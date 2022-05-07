public class Solution {
    /**
     * @param s: the given string
     * @return: the number of A
     */
    public int countA(String s) {
        int res = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == 'A')
                res++;
        }
        return res;
    }
}