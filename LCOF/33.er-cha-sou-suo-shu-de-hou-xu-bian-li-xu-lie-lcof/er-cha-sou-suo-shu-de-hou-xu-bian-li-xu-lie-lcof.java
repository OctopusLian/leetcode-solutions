class Solution {
    public boolean verifyPostorder(int[] postorder) {
        // 解法1，递归
        return recur(postorder, 0, postorder.length - 1);
    }

    boolean recur(int[] postorder, int i, int j) {
        if (i >= j)
            return true;
        int p = i;
        while (postorder[p] < postorder[j])
            p++;
        int m = p;
        while (postorder[p] > postorder[j])
            p++;
        return p == j && recur(postorder, i, m - 1) && recur(postorder, m, j - 1);
    }
}