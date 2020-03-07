import queue
class MaxQueue:

    def __init__(self):
        self.deque = queue.deque()  #创建双向队列

    def max_value(self) -> int:
        return max(self.deque) if self.deque else -1

    def push_back(self, value: int) -> None:
        self.deque.append(value)  #往双向队列里塞入value

    def pop_front(self) -> int:
        return self.deque.popleft() if self.deque else -1 #queue.deque.popleft 获取最左边一个元素，并在队列中删除


# Your MaxQueue object will be instantiated and called as such:
# obj = MaxQueue()
# param_1 = obj.max_value()
# obj.push_back(value)
# param_3 = obj.pop_front()
