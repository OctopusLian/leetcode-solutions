/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode dummy(0);
        ListNode *curr = &dummy;
        int carry = 0;

        while ((l1 != NULL) || (l2 != NULL) || (carry != 0)) {
            int l1_val = (l1 != NULL) ? l1->val : 0;
            int l2_val = (l2 != NULL) ? l2->val : 0;
            int sum = carry + l1_val + l2_val;
            carry = sum / 10;
            curr->next = new ListNode(sum % 10);

            curr = curr->next;
            if (l1 != NULL) l1 = l1->next;
            if (l2 != NULL) l2 = l2->next;
        }

        return dummy.next;
    }
};
