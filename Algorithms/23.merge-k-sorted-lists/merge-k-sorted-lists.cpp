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
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        if (lists.empty()) {
            return NULL;
        }

        ListNode *dummy = new ListNode(INT_MAX);
        ListNode *last = dummy;

        while (true) {
            int count = 0;
	        int index = -1, tempVal = INT_MAX;
            for (int i = 0; i != lists.size(); ++i) {
		        if (NULL == lists[i]) {
		            ++count;
                    if (count == lists.size()) {
                        last->next = NULL;
                        return dummy->next;
                    }
		            continue;
		        }

                // choose the min value in non-NULL ListNode
                if (NULL != lists[i] && lists[i]->val <= tempVal) {
                    tempVal = lists[i]->val;
                    index = i;
                }
            }

	        last->next = lists[index];
	        last = last->next;
	        lists[index] = lists[index]->next;
        }
    }
};
