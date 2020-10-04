/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function(l1, l2) {
    var dfs = l => (l.next && dfs(l.next) || '') + l.val
    return String((BigInt(dfs(l1)) + BigInt(dfs(l2)))).split('').reduce((next,val)=>{return {val, next}}, null)
};