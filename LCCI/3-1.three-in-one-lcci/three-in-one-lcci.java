class TripleInOne {
    private int[] arr;
    private int[] stackTop; // 每个栈当前可push的索引（对应到arr中的索引）
    private int stackSize;

    public TripleInOne(int stackSize) {
        this.stackSize = stackSize;
        arr = new int[stackSize * 3];
        stackTop = new int[]{0, 1, 2};
    }
    
    public void push(int stackNum, int value) {
        int curStackTop = stackTop[stackNum];
        if (curStackTop / 3 == stackSize) {
            // 栈已满
            return;
        }

        arr[curStackTop] = value;
        stackTop[stackNum] += 3;
    }
    
    public int pop(int stackNum) {
        if (stackTop[stackNum] < 3) {
            return -1;
        }
        int value = arr[stackTop[stackNum] - 3];
        stackTop[stackNum] -= 3;
        return value;
    }
    
    public int peek(int stackNum) {
        if (stackTop[stackNum] < 3) {
            return -1;
        }
        return arr[stackTop[stackNum] - 3];
    }
    
    public boolean isEmpty(int stackNum) {
        return stackTop[stackNum] < 3;
    }
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * TripleInOne obj = new TripleInOne(stackSize);
 * obj.push(stackNum,value);
 * int param_2 = obj.pop(stackNum);
 * int param_3 = obj.peek(stackNum);
 * boolean param_4 = obj.isEmpty(stackNum);
 */
