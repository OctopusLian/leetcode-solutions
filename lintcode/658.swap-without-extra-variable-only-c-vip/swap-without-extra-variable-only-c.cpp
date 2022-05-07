class Solution {
public:
    /*
     * @param x: An integer
     * @param y: An integer
     * @return: nothing
     */
    void swap(int& x, int& y) {
        // write your code here
        x = x ^ y;
        y = x ^ y;
        x = x ^ y;
    }
};