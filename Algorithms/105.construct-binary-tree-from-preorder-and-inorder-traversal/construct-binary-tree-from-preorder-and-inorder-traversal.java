/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        if (preorder == null || inorder == null) return null;
        if (preorder.length == 0 || inorder.length == 0) return null;
        if (preorder.length != inorder.length) return null;

        TreeNode root = helper(preorder, 0, preorder.length - 1,
                               inorder, 0, inorder.length - 1);
        return root;
    }
    private TreeNode helper(int[] preorder, int prestart, int preend,
                            int[] inorder, int instart, int inend) {
        // corner cases
        if (prestart > preend || instart > inend) return null;
        // build root TreeNode
        int root_val = preorder[prestart];
        TreeNode root = new TreeNode(root_val);
        // find index of root_val in inorder[]
        int index = findIndex(inorder, instart, inend, root_val);
        // build left subtree
        root.left = helper(preorder, prestart + 1, prestart + index - instart,
               inorder, instart, index - 1);
        // build right subtree
        root.right = helper(preorder, prestart + index - instart + 1, preend,
               inorder, index + 1, inend);
        return root;
    }

    private int findIndex(int[] nums, int start, int end, int target) {
        for (int i = start; i <= end; i++) {
            if (nums[i] == target) return i;
        }
        return -1;
    }
}
