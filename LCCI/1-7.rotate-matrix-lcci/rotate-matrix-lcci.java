class Solution {
    public void rotate(int[][] matrix) {
        int n = matrix.length;
        for (int layer = 0;layer < n / 2;layer++) {
            int first = layer;
            int last = n - 1 - layer;
            for(int i = first;i < last;i++) {
                int offset = i - first;
                int top = matrix[first][i];  //存储上边

                //左边移到右边
                matrix[first][i] = matrix[last-offset][first];
                //下标移到左边
                matrix[last-offset][first] = matrix[last][last - offset];
                //右边移到下边
                matrix[last][last - offset] = matrix[i][last];
                //上边移到右边
                matrix[i][last] = top;
            }
        }
        
    }
}
