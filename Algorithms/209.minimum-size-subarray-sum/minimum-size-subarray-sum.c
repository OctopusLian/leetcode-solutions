int minSubArrayLen(int s, int* nums, int numsSize) {
    //解法1，暴力法
    if (numsSize == 0) {
        return 0;
    }
    int ans = INT_MAX;
    for (int i = 0; i < numsSize; i++) {
        int sum = 0;
        for (int j = i; j < numsSize; j++) {
            sum += nums[j];
            if (sum >= s) {
                ans = fmin(ans, j - i + 1);
                break;
            }
        }
    }
    return ans == INT_MAX ? 0 : ans;
}