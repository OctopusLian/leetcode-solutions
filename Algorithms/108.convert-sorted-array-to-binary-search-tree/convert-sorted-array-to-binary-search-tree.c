/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */


struct TreeNode* helper(int* nums, int left, int right) {
    if (left > right) {
        return NULL;
    }

    // 总是选择中间位置左边的数字作为根节点
    int mid = (left + right) / 2;

    struct TreeNode* root = (struct TreeNode*)malloc(sizeof(struct TreeNode));
    root->val = nums[mid];
    root->left = helper(nums, left, mid - 1);
    root->right = helper(nums, mid + 1, right);
    return root;
}

struct TreeNode* sortedArrayToBST(int* nums, int numsSize) {
    //解法1，中序遍历，总是选择中间位置左边的数字作为根节点
    return helper(nums, 0, numsSize - 1);
}