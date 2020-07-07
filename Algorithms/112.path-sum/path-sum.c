/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */


bool hasPathSum(struct TreeNode* root, int sum){
    //解法1，递归
    if (root == NULL) {
        return false;
    }
    if (root->left == NULL && root->right == NULL) {
        return sum == root->val;
    }
    return hasPathSum(root->left,sum - root->val) || hasPathSum(root->right,sum - root->val);
}