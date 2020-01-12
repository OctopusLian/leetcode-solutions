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

        RandomListNode *dummy = new RandomListNode(0);
        RandomListNode *curNode = dummy;
        unordered_map<RandomListNode *, RandomListNode *> randomMap;
        while(head != NULL) {
            // link newNode to new List
            RandomListNode *newNode = new RandomListNode(head->label);
            curNode->next = newNode;
            // map old node head to newNode
            randomMap[head] = newNode;
            // copy old node random pointer
            newNode->random = head->random;

            head = head->next;
            curNode = curNode->next;
        }

        // re-mapping old random node to new node
        curNode = dummy->next;
        while (curNode != NULL) {
            if (curNode->random != NULL) {
                curNode->random = randomMap[curNode->random];
            }
            curNode = curNode->next;
        }

        return dummy->next;
    }
};


//有Bug
