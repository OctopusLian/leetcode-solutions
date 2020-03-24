class Solution {
    public int massage(int[] nums) {
	//迭代法
        int[] memo = new int[nums.length + 2];
        memo[nums.length] = 0;
        memo[nums.length + 1] = 0;
        for (int i = nums.length - 1;i >=0;i--) {
            int bestWith = nums[i] + memo[i + 2];
            int bestWithout = memo[i + 1];
            memo[i] = Math.max(bestWith,bestWithout);
        }
        return memo[0];
    }
}
