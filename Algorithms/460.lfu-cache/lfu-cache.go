type LFUCache struct {
  cache map[int]*Node
  freq map[int]*DoubleList
  ncap, size, minFreq int
}

func Constructor(capacity int) LFUCache {
  return LFUCache {
    cache: make(map[int]*Node),
    freq: make(map[int]*DoubleList),
    ncap: capacity,
  }
}

func (this *LFUCache) Get(key int) int {
  if node, ok := this.cache[key]; ok {
    this.IncFreq(node)
    return node.val
  }
  return -1
}

func (this *LFUCache) Put(key, value int) {
  if this.ncap == 0 {
    return
  }
  if node, ok := this.cache[key]; ok {
    node.val = value
    this.IncFreq(node)
  } else {
    if this.size >= this.ncap {
      node := this.freq[this.minFreq].RemoveLast()
      delete(this.cache, node.key)
      this.size--
    }
    x := &Node{key: key, val: value, freq: 1}
    this.cache[key] = x
    if this.freq[1] == nil {
      this.freq[1] = CreateDL()
    }
    this.freq[1].AddFirst(x)
    this.minFreq = 1
    this.size++
  }
}

func (this *LFUCache) IncFreq(node *Node) {
  _freq := node.freq
  this.freq[_freq].Remove(node)
  if this.minFreq == _freq && this.freq[_freq].IsEmpty() {
    this.minFreq++
    delete(this.freq, _freq)
  }

  node.freq++
  if this.freq[node.freq] == nil {
    this.freq[node.freq] = CreateDL()
  }
  this.freq[node.freq].AddFirst(node)
}

type DoubleList struct {
  head, tail *Node
}

type Node struct {
  prev, next *Node
  key, val, freq int
}

func CreateDL() *DoubleList {
  head, tail := &Node{}, &Node{}
  head.next, tail.prev = tail, head
  return &DoubleList {
    head: head,
    tail: tail,
  }
}

func (this *DoubleList) AddFirst(node *Node) {
  node.next = this.head.next
  node.prev = this.head

  this.head.next.prev = node
  this.head.next = node
}

func (this *DoubleList) Remove(node *Node) {
  node.prev.next = node.next
  node.next.prev = node.prev

  node.next = nil
  node.prev = nil
}

func (this *DoubleList) RemoveLast() *Node {
  if this.IsEmpty() {
    return nil
  }

  last := this.tail.prev
  this.Remove(last)

  return last
}

func (this *DoubleList) IsEmpty() bool {
  return this.head.next == this.tail
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
