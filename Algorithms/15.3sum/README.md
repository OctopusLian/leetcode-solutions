## 题解1 - 排序 + 哈希表 + 2 Sum  

相比之前的 [2 Sum](http://algorithm.yuanbin.me/zh-hans/integer_array/2_sum.html), 3 Sum 又多加了一个数，按照之前 2 Sum 的分解为『1 Sum + 1 Sum』的思路，我们同样可以将 3 Sum 分解为『1 Sum + 2 Sum』的问题，具体就是首先对原数组排序，排序后选出第一个元素，随后在剩下的元素中使用 2 Sum 的解法。

### Python  

```python
class Solution:
    """
    @param numbersbers : Give an array numbersbers of n integer
    @return : Find all unique triplets in the array which gives the sum of zero.
    """
    def threeSum(self, numbers):
        triplets = []
        length = len(numbers)
        if length < 3:
            return triplets

        numbers.sort()
        for i in xrange(length):
            target = 0 - numbers[i]
            # 2 Sum
            hashmap = {}
            for j in xrange(i + 1, length):
                item_j = numbers[j]
                if (target - item_j) in hashmap:
                    triplet = [numbers[i], target - item_j, item_j]
                    if triplet not in triplets:
                        triplets.append(triplet)
                else:
                    hashmap[item_j] = j

        return triplets
```

### 源码分析  

1. 异常处理，对长度小于3的直接返回。
2. 排序输入数组，有助于提高效率和返回有序列表。
3. 循环遍历排序后数组，先取出一个元素，随后求得 2 Sum 中需要的目标数。
4. 由于本题中最后返回结果不能重复，在加入到最终返回值之前查重。

由于排序后的元素已经按照大小顺序排列，且在2 Sum 中先遍历的元素较小，所以无需对列表内元素再排序。

### 复杂度分析  

排序时间复杂度 `$$O(n \log n)$$`, 两重`for`循环，时间复杂度近似为 `$$O(n^2)$$`，使用哈希表(字典)实现，空间复杂度为 `$$O(n)$$`.

目前这段源码为比较简易的实现，leetcode 上的运行时间为500 + ms, 还有较大的优化空间，嗯，后续再进行优化。

### C++ 
```c++
class Solution {
public:
    vector<vector<int> > threeSum(vector<int> &num) 
    {
        vector<vector<int> > result;
        if (num.size() < 3) return result;
        
        int ans = 0;

        sort(num.begin(), num.end());
        
        for (int i = 0;i < num.size() - 2; ++i)
        {
            if (i > 0 && num[i] == num[i - 1])  
                continue;
            int j = i + 1;
            int k = num.size() - 1;

            while (j < k)
            {
                ans = num[i] + num[j] + num[k];

                if (ans == 0)
                {
                    result.push_back({num[i], num[j], num[k]});
                    ++j;
                    while (j < num.size() && num[j] == num[j - 1])
                        ++j;
                    --k;
                    while (k >= 0 && num[k] == num[k + 1])
                        --k;
                }
                else if (ans > 0) 
                    --k;
                else 
                    ++j;
            }
        }
        
        return result;
    }
};
```
###源码分析  

同python解法不同，没有使用hash map
```
S = {-1 0 1 2 -1 -4}
排序后：
S = {-4 -1 -1 0 1 2}
      ↑  ↑        ↑
      i  j        k
         →        ←
i每轮只走一步，j和k根据S[i]+S[j]+S[k]=ans和0的关系进行移动，且j只向后走（即S[j]只增大），k只向前走（即S[k]只减小）
如果ans>0说明S[k]过大，k向前移；如果ans<0说明S[j]过小，j向后移；ans==0即为所求。
至于如何取到所有解，看代码即可理解，不再赘述。
```
###复杂度分析  

外循环i走了n轮,每轮j和k一共走n-i步，所以时间复杂度为`$$O(n^2)$$`。
最终运行时间为52ms
## Reference

- [3Sum | 九章算法](http://www.jiuzhang.com/solutions/3sum/)
- [A simply Python version based on 2sum - O(n^2) - Leetcode Discuss](https://leetcode.com/discuss/32455/a-simply-python-version-based-on-2sum-o-n-2)

