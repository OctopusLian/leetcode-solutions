class Solution {
public:
    vector<int> distributeCandies(int candies, int num_people) {
        vector<int> ans(num_people,0);
        int i = 0;
        while (candies != 0) {  //当糖果数不为0时，继续这个循环
            // i % num_people作用为转为起始，从第一个小朋友开始
            ans[i % num_people] += min(candies, i + 1);
            candies -= min(candies, i + 1);
            ++i;
        }
        return ans;
    }
};
