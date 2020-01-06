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
    ListNode* reverseBetween(ListNode* head, int m, int n) {
        if (head == NULL || m > n) {
            return NULL;
        }

        ListNode *dummy = new ListNode(0);
        dummy->next = head;
        ListNode *node = dummy;

        for (int i = 1; i != m; ++i) {
            if (node == NULL) {
                return NULL;
            } else {
                node = node->next;
            }
        }

        ListNode *premNode = node;
        ListNode *mNode = node->next;
        ListNode *nNode = mNode, *postnNode = nNode->next;
        for (int i = m; i != n; ++i) {
            if (postnNode == NULL) {
                return NULL;
            }

            ListNode *temp = postnNode->next;
            postnNode->next = nNode;
            nNode = postnNode;
            postnNode = temp;
        }
        premNode->next = nNode;
        mNode->next = postnNode;

        return dummy->next;
    }
};
