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
    ListNode* deleteNode(ListNode* head, int val) {
        struct ListNode** indirect = &head; 
    //struct ListNode* temp;
    for(; *indirect; indirect = &((*indirect)->next)){
        if((*indirect)->val == val){
            //temp = *indirect;
            *indirect = (*indirect)->next;
            //free(temp);
            break;
        }
    }
    return head;
    }
};


/*
Linus Solution
http://wordaligned.org/articles/two-star-programming 
*/
