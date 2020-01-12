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
    ListNode* sortList(ListNode* head) {
        int len_list = 0;
        ListNode *p=head;
        while(p){
            p = p->next;
            len_list++;
        }
        ListNode *l_list,*r_list,**p_merge_list;
        for(int i = 1; i < len_list; i <<= 1){
            r_list = l_list = head;
            p_merge_list = &head;
            for(int j = 0; j < len_list - i ; j += i << 1){

                for(int k = 0; k < i; ++k) r_list=r_list->next;
                int l_len=i,r_len=min(i, len_list - j - i);

                while(l_len || r_len ){
                    if(r_len > 0 && (l_len == 0 || r_list->val <= l_list->val)){
                        *p_merge_list = r_list;
                        p_merge_list=&(r_list->next);
                        r_list = r_list->next;
                        --r_len;
                    }
                    else{
                        *p_merge_list = l_list;
                        p_merge_list=&(l_list->next);
                        l_list = l_list->next;

                        --l_len;
                    }
                }
                l_list=r_list;
            }
            *p_merge_list = r_list;

        }
        return head;
    }
};
