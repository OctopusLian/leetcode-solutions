## 题解

此题出自 *CTCI* 题 2.4，依据题意，是要根据值x对链表进行分割操作，具体是指将所有小于x的节点放到不小于x的节点之前，咋一看和快速排序的分割有些类似，但是这个题的不同之处在于只要求将小于x的节点放到前面，而并不要求对元素进行排序。

这种分割的题使用两路指针即可轻松解决。左边指针指向小于x的节点，右边指针指向不小于x的节点。由于左右头节点不确定，我们可以使用两个dummy节点。  

### 源码分析

1. 异常处理
2. 引入左右两个dummy节点及left和right左右尾指针
3. 遍历原链表
4. 处理右链表，置`right->next`为空(否则如果不为尾节点则会报错，处理链表时 以 null 为判断)，将右链表的头部链接到左链表尾指针的next，返回左链表的头部

### 复杂度分析

遍历链表一次，时间复杂度近似为 `$$O(n)$$`, 使用了两个 dummy 节点及中间变量，空间复杂度近似为 `$$O(1)$$`.


### 源码分析

1. 迭代能正常进行的条件为`(NULL != l1) || (NULL != l2) || (0 != carry)`, 缺一不可。
2. 对于空指针节点的处理可以用相对优雅的方式处理 - `int l1_val = (NULL == l1) ? 0 : l1->val;`
3. ~~生成新节点时需要先判断迭代终止条件 - `(NULL == l1) && (NULL == l2) && (0 == carry)`, 避免多生成一位数0。~~ 使用 dummy 节点可避免这一情况。

### 复杂度分析

没啥好分析的，时间和空间复杂度均为 `$$O(max(L1, L2))$$`.

## Reference

- *CC150 Chapter 9.2* 题2.5，中文版 p123
- [Add two numbers represented by linked lists | Set 1 - GeeksforGeeks](http://www.geeksforgeeks.org/add-two-numbers-represented-by-linked-lists/)
# Two Lists Sum Advanced <i class="fa fa-star"></i><i class="fa fa-star"></i><i class="fa fa-star"></i>

