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
        RandomListNode *dummy = new RandomListNode(0);
        RandomListNode *curNode = dummy;
        unordered_map<RandomListNode *, RandomListNode *> hash_map;
        while(head != NULL) {
            // link newNode to new List
            RandomListNode *newNode = NULL;
            if (hash_map.count(head) > 0) {
                newNode = hash_map[head];
            } else {
                newNode = new RandomListNode(head->label);
                hash_map[head] = newNode;
            }
            curNode->next = newNode;
            // re-mapping old random node to new node
            if (head->random != NULL) {
                if (hash_map.count(head->random) > 0) {
                    newNode->random = hash_map[head->random];
                } else {
                    newNode->random = new RandomListNode(head->random->label);
                    hash_map[head->random] = newNode->random;
                }
            }

            head = head->next;
            curNode = curNode->next;
        }

        return dummy->next;
    }
};
