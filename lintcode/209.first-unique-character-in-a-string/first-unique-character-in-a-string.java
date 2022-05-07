public class Solution {
    /**
     * @param str: str: the given string
     * @return: char: the first unique character in a given string
     */
    public char firstUniqChar(String str) {

        if (str == null) {
            return '0';
        }

        String[] s = new String[str.length()];
        for (int i = 0; i < str.length(); i++) {
            // 将字符串转成字符串数组
            s[i] = str.substring(i, i + 1);
        }

        char a = 'a';
        for (int i = 0; i < str.length(); i++) {
            if (str.indexOf(s[i]) == str.lastIndexOf(s[i])) {
                // 获取字符出现的下标
                a = s[i].charAt(0);
                break;
            }
        }

        return a;
    }
}