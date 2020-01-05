## 题解

题 [Majority Number](http://algorithm.yuanbin.me/zh-hans/math_and_bit_manipulation/majority_number.html) 的升级版，之前那道题是『两两抵消』，这道题自然则需要『三三抵消』，不过『三三抵消』需要注意不少细节，比如两个不同数的添加顺序和添加条件。  


### 源码分析

首先处理`count == 0`的情况，这里需要注意的是`count2 == 0 && key1 = num`, 不重不漏。最后再次遍历原数组也必不可少，因为由于添加顺序的区别，count1 和 count2的大小只具有相对意义，还需要最后再次比较其真实计数器值。

### 复杂度分析

时间复杂度 `$$O(n)$$`, 空间复杂度 `$$O(2 \times 2) = O(1)$$`.

## Reference

- [Majority Number II 参考程序 Java/C++/Python](http://www.jiuzhang.com/solutions/majority-number-ii/)
