class Solution {
    public boolean searchMatrix(int[][] matrix, int target) {
        if (matrix == null || matrix.length == 0 || matrix[0] == null) {
            return false;
        }

        int ROW = matrix.length, COL = matrix[0].length;
        int lb = -1, ub = ROW * COL;
        while (lb + 1 < ub) {
            int mid = lb + (ub - lb) / 2;
            if (matrix[mid / COL][mid % COL] < target) {
                lb = mid;
            } else {
                if (matrix[mid / COL][mid % COL] == target) {
                    return true;
                }
                ub = mid;
            }
        }

        return false;
    }
}
