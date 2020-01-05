## 题解

遍历之，遇到当前节点和下一节点的值相同时，删除下一节点，并将当前节点`next`值指向下一个节点的`next`, 当前节点首先保持不变，直到相邻节点的值不等时才移动到下一节点。  


### 源码分析

1. 首先进行异常处理，判断head是否为NULL
2. 遍历链表，`curr->val == curr->next->val`时，保存`curr->next`，便于后面释放内存(非C/C++无需手动管理内存)
3. 不相等时移动当前节点至下一节点，注意这个步骤必须包含在`else`中，否则逻辑较为复杂

~~`while` 循环处也可使用`curr != null && curr.next != null`, 这样就不用单独判断`head` 是否为空了，但是这样会降低遍历的效率，因为需要判断两处。~~使用双重`while`循环可只在内循环处判断，避免了冗余的判断，谢谢 @xuewei4d 提供的思路。

### 复杂度分析

遍历链表一次，时间复杂度为 `$$O(n)$$`, 使用了一个中间变量进行遍历，空间复杂度为 `$$O(1)$$`.

## Reference

- [Remove Duplicates from Sorted List 参考程序 | 九章](http://www.jiuzhang.com/solutions/remove-duplicates-from-sorted-list/)
