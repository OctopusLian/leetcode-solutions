# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

# 二叉树
# f(root) = f(root->left) + f(root->right)

class Solution:
    def allPossibleFBT(self, N: int) -> List[TreeNode]:
        if N == 1:
            return [TreeNode(0)]
        res = []
        for i in range(1,N,2):
            nLeft = i
            nRight = N -1 - i
            for left in self.allPossibleFBT(nLeft):
                for right in self.allPossibleFBT(nRight):
                    root = TreeNode(0)
                    root.left = left
                    root.right = right
                    res.append(root)
        return res