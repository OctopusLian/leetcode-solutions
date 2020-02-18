class Solution {
    public int reverseBits(int num) {
        if (~num == 0) return Integer.BYTES * 8;

        int currentLength = 0;
        int previousLength = 0;
        int maxLength = 1;
        while (num != 0) {
            if ((num & 1) == 1) {
                currentLength++;
            } else if ((num & 1) == 0) {
                previousLength = (num & 2) == 0 ? 0 : currentLength;
                currentLength = 0;
            }
            maxLength = Math.max(previousLength + currentLength + 1,maxLength);
            num >>>= 1;
        }
        return maxLength;
    }
}
