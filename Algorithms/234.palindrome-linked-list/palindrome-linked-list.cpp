class Solution {
public:
    bool isPalindrome(ListNode* head) {
        if (!head || !head->next) return true;  

        // find middle
        ListNode* slow = head, *fast = head;
        while (fast && fast->next) {               
            slow = slow->next;
            fast = fast->next->next;
        }

        // skip mid node if the number of ListNode is odd
        if (fast) slow = slow->next;    

        // reverse right part of List
        ListNode* rHead = reverse(slow);  
        ListNode* lCurr = head, *rCurr = rHead;
        while (rCurr) {
            if (rCurr->val != lCurr->val) {
                reverse(rHead);
                return false;
            }
            lCurr = lCurr->next;
            rCurr = rCurr->next;
        }
        // recover right part of List
        reverse(rHead);            

        return true;
    }

    ListNode* reverse(ListNode* head) {
        ListNode* prev = NULL;
        while (head) {                           
            ListNode* after = head->next;   
            head->next = prev;
            prev = head;
            head = after;
        }
        return prev;
    }
}
