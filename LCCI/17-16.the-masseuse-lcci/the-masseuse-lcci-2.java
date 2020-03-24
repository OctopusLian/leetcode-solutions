class Solution {
    public int massage(int[] nums) {
	//优化时间和空间的迭代
        int oneAway = 0;
        int twoAway = 0;
        for (int i = nums.length - 1;i >= 0;i--) {
            int bestWith = nums[i] + twoAway;
            int bestWithout = oneAway;
            int current = Math.max(bestWith,bestWithout);
            twoAway = oneAway;
            oneAway = current;
        }
        return oneAway;
    }
}
