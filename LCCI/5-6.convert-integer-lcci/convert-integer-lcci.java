class Solution {
    public int convertInteger(int A, int B) {
        int count = 0;
        for (int c = A ^ B;c != 0;c = c & (c-1)) {
            count++;
        }
        return count;
    }
}
