## 题解

此题在上题的基础上加了位置要求，只翻转指定区域的链表。由于链表头节点不确定，祭出我们的dummy杀器。此题边界条件处理特别tricky，需要特别注意。

1. 由于只翻转指定区域，分析受影响的区域为第m-1个和第n+1个节点
2. 找到第m个节点，使用for循环n-m次，使用上题中的链表翻转方法
3. 处理第m-1个和第n+1个节点
4. 返回dummy->next  


### 源码分析

1. 处理异常
2. 使用dummy辅助节点
3. 找到premNode——m节点之前的一个节点
4. 以nNode和postnNode进行遍历翻转，注意考虑在遍历到n之前postnNode可能为空
5. 连接premNode和nNode，`premNode->next = nNode;`
6. 连接mNode和postnNode，`mNode->next = postnNode;`

**务必注意node 和node->next的区别！！**，node指代节点，而`node->next`指代节点的下一连接。
