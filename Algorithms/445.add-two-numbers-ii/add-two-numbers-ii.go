/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  var a,b []int
  for p:=l1;p!=nil;p=p.Next  {
    a = append(a,p.Val)

  }
  for p:=l2;p!=nil;p=p.Next {
    b = append(b,p.Val)
  }
  var res *ListNode
  carry := 0
  aLen,bLen := len(a),len(b)
  for i,j:=0,0; i<aLen || j<bLen; i,j=i+1,j+1{
    sum := carry
    if i< aLen{
      sum += a[aLen-i-1]
    }
    if j< bLen{
      sum += b[bLen-j-1]
    }
    node := ListNode{Val:sum%10,Next:nil}
    if res == nil{
      res = &node
    }else{
      node.Next = res
      res = &node
    }
    carry = sum / 10
  }
  if carry != 0{
    res = &ListNode{Val:carry,Next:res}
  }
  return res
}
