## 题解1 - 非递归  

联想到同样也可能需要翻转的数组，在数组中由于可以利用下标随机访问，翻转时使用下标即可完成。而在单向链表中，仅仅只知道头节点，而且只能单向往前走，故需另寻出路。分析由`1->2->3`变为`3->2->1`的过程，由于是单向链表，故只能由1开始遍历，1和2最开始的位置是`1->2`，最后变为`2->1`，故从这里开始寻找突破口，探讨如何交换1和2的节点。

```
temp = head->next;
head->next = prev;
prev = head;
head = temp;
```

要点在于维护两个指针变量`prev`和`head`, 翻转相邻两个节点之前保存下一节点的值，分析如下图所示：

![Reverse Linked List](https://raw.githubusercontent.com/billryan/algorithm-exercise/master/shared-files/images/reverse_linked_list_i.jpg)

1. 保存head下一节点
2. 将head所指向的下一节点改为prev
3. 将prev替换为head，波浪式前进
4. 将第一步保存的下一节点替换为head，用于下一次循环  


## 题解2 - 递归  

递归的终止步分三种情况讨论：

1. 原链表为空，直接返回空链表即可。
2. 原链表仅有一个元素，返回该元素。
3. 原链表有两个以上元素，由于是单链表，故翻转需要自尾部向首部逆推。

由尾部向首部逆推时大致步骤为先翻转当前节点和下一节点，然后将当前节点指向的下一节点置空(否则会出现死循环和新生成的链表尾节点不指向空)，如此递归到头节点为止。新链表的头节点在整个递归过程中一直没有变化，逐层向上返回。  


## Reference

- [全面分析再动手的习惯：链表的反转问题（递归和非递归方式） - 木棉和木槿 - 博客园](http://www.cnblogs.com/kubixuesheng/p/4394509.html)
- [data structures - Reversing a linked list in Java, recursively - Stack Overflow](http://stackoverflow.com/questions/354875/reversing-a-linked-list-in-java-recursively)
- [反转单向链表的四种实现（递归与非递归，C++） | 宁心勉学，慎思笃行](http://ceeji.net/blog/reserve-linked-list-cpp/)
- [iteratively and recursively Java Solution - Leetcode Discuss](https://leetcode.com/discuss/37804/iteratively-and-recursively-java-solution)
