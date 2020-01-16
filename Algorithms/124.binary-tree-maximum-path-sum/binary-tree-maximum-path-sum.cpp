/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    int maxPathSum(TreeNode* root) {
        if (NULL == root) {
            return 0;
        }

        int result = INT_MIN;
        stack<TreeNode *> s;
        s.push(root);
        while (!s.empty()) {
            TreeNode *node = s.top();
            s.pop();

            int temp_path_sum = node->val + singlePathSum(node->left) \
                                          + singlePathSum(node->right);

            if (temp_path_sum > result) {
                result = temp_path_sum;
            }

            if (NULL != node->right) s.push(node->right);
            if (NULL != node->left) s.push(node->left);
        }

        return result;
    }
private:
    int singlePathSum(TreeNode *root) {
        if (NULL == root) {
            return 0;
        }

        int path_sum = max(singlePathSum(root->left), singlePathSum(root->right));
        return max(0, (root->val + path_sum));
    }
};
