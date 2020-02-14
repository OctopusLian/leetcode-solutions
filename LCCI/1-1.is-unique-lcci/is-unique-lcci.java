class Solution {
    public boolean isUnique(String astr) {
        if (astr.length() > 128) return false;

        boolean[] char_set = new boolean[128];
        for (int i = 0;i < astr.length();i++) {
            int val = astr.charAt(i);
            if (char_set[val]) {
                return false;
            }
            char_set[val] = true;
        }
        return true;
    }
}
