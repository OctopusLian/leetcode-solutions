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
    void reorderList(ListNode* head) {
        if (head == NULL || head->next == NULL) return;

        // find middle
        ListNode *slow = head, *fast = head->next;
        while (fast != NULL && fast->next != NULL) {
            slow = slow->next;
            fast = fast->next->next;
        }
        ListNode *rHead = slow->next;
        slow->next = NULL;

        // reverse ListNode on the right side
        ListNode *prev = NULL;
        while (rHead != NULL) {
            ListNode *temp = rHead->next;
            rHead->next = prev;
            prev = rHead;
            rHead = temp;
        }

        // merge two list
        rHead = prev;
        ListNode *lHead = head;
        while (lHead != NULL && rHead != NULL) {
            ListNode *temp1 = lHead->next;
            lHead->next = rHead;
            ListNode *temp2 = rHead->next;
            rHead->next = temp1;
            lHead = temp1;
            rHead = temp2;
        }
    }
};
