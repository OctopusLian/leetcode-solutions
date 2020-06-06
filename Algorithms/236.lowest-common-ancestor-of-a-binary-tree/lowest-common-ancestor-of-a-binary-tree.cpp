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
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if (!root) return root;

        if (root == p || root == q) return root;

        TreeNode* l = lowestCommonAncestor(root->left,p,q);
        TreeNode* r = lowestCommonAncestor(root->right,p,q);

        if (l && r)
            return root;
        
        if(l) return l;
        else return r;
    }
};

//1，path1(root -> p)，part2(root->q)；最后分叉点，就是公共祖先
//2，递归来查询：左右字树寻找p和q
//only left -> left子树
//only left -> left子树