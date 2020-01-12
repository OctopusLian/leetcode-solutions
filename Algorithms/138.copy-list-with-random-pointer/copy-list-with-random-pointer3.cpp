/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* next;
    Node* random;
    
    Node(int _val) {
        val = _val;
        next = NULL;
        random = NULL;
    }
};
*/
class Solution {
public:
    Node* copyRandomList(Node* head) {
        if (head == NULL) return NULL;

        RandomListNode *curr = head;
        // step1: generate new List with node
        while (curr != NULL) {
            RandomListNode *newNode = new RandomListNode(curr->label);
            newNode->next = curr->next;
            curr->next = newNode;
            //
            curr = curr->next->next;
        }
        // step2: copy random
        curr = head;
        while (curr != NULL) {
            if (curr->random != NULL) {
                curr->next->random = curr->random->next;
            }
            curr = curr->next->next;
        }
        // step3: split original and new List
        RandomListNode *newHead = head->next;
        curr = head;
        while (curr != NULL) {
            RandomListNode *newNode = curr->next;
            curr->next = curr->next->next;
            curr = curr->next;
            if (newNode->next != NULL) {
                newNode->next = newNode->next->next;
            }
        }

        return newHead;
    }
};
