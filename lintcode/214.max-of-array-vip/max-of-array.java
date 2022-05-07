public class Solution {
    /**
     * @param A: An integer
     * @return: a float number
     */
    public float maxOfArray(float[] A) {
        float max = A[0];
        for (int i = 1; i < A.length; i++) {
            if (A[i] > max) {
                max = A[i];
            }
        }
        return max;
    }
}