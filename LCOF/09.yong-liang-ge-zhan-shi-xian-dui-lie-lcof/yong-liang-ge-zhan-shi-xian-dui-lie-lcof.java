class CQueue {
    //创建栈 stack1 用来充当队列
    Stack<Integer> stack1;
    //创建栈 stack2 用来辅助 stack1 执行队列的一些复杂操作
    Stack<Integer> stack2;
    //初始化队列
    public CQueue() {
        //初始化 stack1
        stack1 = new Stack<Integer>();

        //初始化 stack2
        stack2 = new Stack<Integer>();
    }
    
    //在队列的尾部添加元素
    public void appendTail(int value) {
        /**
        直接将元素放到 stack1 中
        比如原队列
        ---------------------
        队尾 1 2 3 4 5 队头
        --------------------
        当前 value 为 6
        那么由 stack1 和 stack2 组成的队列就是
        --------------------
        队尾 6 1 2 3 4 5 队头
        --------------------
        */
        stack1.push(value);
    }
    
    //在队列的头部删除元素
    public int deleteHead() {
        //1，如果 stack2 栈不为空，说明 stack2 里面已经存储了一些元素，并且stack 的栈顶元素就是两个栈中最早加入的元素
        if (!stack2.isEmpty()) {
            //返回 stack2 的栈顶元素，满足了队列先进先出的特点
            return stack2.pop();
        }

        //2，如果 stack2 为空，并且发现stack1 也为空，说明 stack1和stack2 构建的队列中没有元素
        if (stack1.isEmpty()) {
            //直接返回-1
            return -1;
        }

        //3，如果stack2为空，但stack1不为空，那么需要将stack1中的元素依次 【倒序】 放入stack2中，
        //对于stack1 来说，越早加入的元素在 【栈顶】，越晚加入的元素在【栈顶】
        //由于队列是【先进先出】，所以删除的应该是stack1的【栈底】元素
        while (!stack1.isEmpty()) {
            //获取stack1的栈顶元素并将该元素从stack1弹出
            int topValue = stack1.pop();
            //把该元素加入到stack2，这样stack2的栈顶元素就是stack1的栈底元素
            stack2.push(topValue);
        }

        //4，返回stack2的栈顶元素，满足了队列先进先出的特点
        return stack2.pop();
    }
}

/**
 * Your CQueue object will be instantiated and called as such:
 * CQueue obj = new CQueue();
 * obj.appendTail(value);
 * int param_2 = obj.deleteHead();
 */