/*
一维数组的坐标变换
1,枚举 left bar x * right bar y, (x-y)*height_diff
*/

class Solution {
    public int maxArea(int[] height) {
        int max = 0;
        for (int i = 0;i < height.length - 1;++i) {
            for (int j = i + 1;j < height.length;++j) {
                int area = (j - i) * Math.min(height[i],height[j]);
                max = Math.max(max,area);
            }
        }
        return max;
    }
}
