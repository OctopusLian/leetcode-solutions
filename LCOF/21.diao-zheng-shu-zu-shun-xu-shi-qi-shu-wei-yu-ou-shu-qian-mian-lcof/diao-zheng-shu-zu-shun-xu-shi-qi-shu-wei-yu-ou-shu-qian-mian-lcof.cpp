class Solution {
public:
    vector<int> exchange(vector<int>& nums) {
        int left = 0, right = nums.size() - 1;  //创建首尾双指针
        while (left < right) {  //如果左指针一直在右指针的左边
            if ((nums[left] & 1) != 0) {  //如果左指针指向的数是奇数
                left ++;
                continue;
            }
            if ((nums[right] & 1) != 1) {  //如果右指针指向的是是偶数
                right --;
                continue;
            }
            swap(nums[left], nums[right]);  //以上条件都不满足，交换两个指针指向的数
        }
        return nums;
    }
};
