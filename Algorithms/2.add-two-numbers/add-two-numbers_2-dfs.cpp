/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        return dfs(l1, l2, 0);
    }

    ListNode* dfs(ListNode* l, ListNode* r, int i) {
        if (!l && !r && !i) return nullptr;
        int sum = (l ? l->val : 0) + (r ? r->val : 0) + i;
        ListNode* node = new ListNode(sum % 10);
        node->next = dfs(l ? l->next : nullptr, r ? r->next : nullptr, sum / 10);
        return node;
    } 
};