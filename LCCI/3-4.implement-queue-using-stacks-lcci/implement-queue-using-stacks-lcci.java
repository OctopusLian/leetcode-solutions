class MyQueue {
    private Stack<Integer> stackNewest,stackOldest;

    /** Initialize your data structure here. */
    public MyQueue() {
        stackNewest = new Stack<>();
        stackOldest = new Stack<>();
    }
    
    /** Push element x to the back of queue. */
    public void push(int x) {
        stackNewest.push(x);
    }
    
    /** Removes the element from in front of queue and returns that element. */
    public int pop() {
        while (!stackNewest.isEmpty()) {
            stackOldest.push(stackNewest.pop());
        }
        int res = stackOldest.pop();
        while (!stackOldest.isEmpty()) {
            stackNewest.push(stackOldest.pop());
        }
        return res;
    }
    
    /** Get the front element. */
    public int peek() {
        while (!stackNewest.isEmpty()) {
            stackOldest.push(stackNewest.pop());
        }
        int res = stackOldest.peek();
        while (!stackOldest.isEmpty()) {
            stackNewest.push(stackOldest.pop());
        }
        return res;
    }
    
    /** Returns whether the queue is empty. */
    public boolean empty() {
        return stackNewest.isEmpty();
    }
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue obj = new MyQueue();
 * obj.push(x);
 * int param_2 = obj.pop();
 * int param_3 = obj.peek();
 * boolean param_4 = obj.empty();
 */
