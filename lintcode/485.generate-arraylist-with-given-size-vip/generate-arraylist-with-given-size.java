public class Solution {
    /**
     * @param size: An integer
     * @return: An integer list
     */
    public List<Integer> generate(int size) {
        ArrayList<Integer> result = new ArrayList<Integer>();
        for (int i = 1; i <= size; ++i)
            result.add(i);
        return result;
    }
}