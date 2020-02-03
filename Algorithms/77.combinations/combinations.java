class Solution {
    public List<List<Integer>> combine(int n, int k) {
        assert(n >= 1 && n >= k && k >= 1);
        
        List<List<Integer>> result = new ArrayList<List<Integer>>();
        List<Integer> list = new ArrayList<Integer>();
        dfs(n, k, 1, list, result);
        
        return result;
    }
    private void dfs(int n, int k, int pos, List<Integer> list,
                     List<List<Integer>> result) {
        
        if (list.size() == k) {
            result.add(new ArrayList<Integer>(list));
            return;
        }
        for (int i = pos; i <= n; i++) {
            list.add(i);
            dfs(n, k, i + 1, list, result);
            list.remove(list.size() - 1);
        }
    }
}
