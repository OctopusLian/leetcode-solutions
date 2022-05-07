/**
 * Definition of TreeNode:
 * public class TreeNode {
 *     public int val;
 *     public TreeNode left, right;
 *     public TreeNode(int val) {
 *         this.val = val;
 *         this.left = this.right = null;
 *     }
 * }
 */

public class Solution {
    /**
     * @param root: a root of binary tree
     * @return: return a integer
     */
    public int diameterOfBinaryTree(TreeNode root) {
        return getDeepth(root.left) + getDeepth(root.right);
    }
    //递归求树的高度
    private int getDeepth(TreeNode root){
        if(root == null){
            return 0;
        }
        return Math.max(getDeepth(root.left),getDeepth(root.right)) + 1;
    }
}