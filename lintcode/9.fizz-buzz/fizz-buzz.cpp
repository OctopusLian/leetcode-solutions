class Solution {
public:
    /**
     * @param n: An integer
     * @return: A list of strings.
     */
    vector<string> fizzBuzz(int n) {
        // write your code here
        vector<string> results;
        for (int i = 1; i <= n; i++) {
            //能同时被3跟5整除的数，肯定也能被3整除，也能被5整除，所以要先判断能否被15整除
            if (i % 15 == 0) {
                results.push_back("fizz buzz");
            } else if (i % 5 == 0) {
                results.push_back("buzz");
            } else if (i % 3 == 0) {
                results.push_back("fizz");
            } else {
                results.push_back(to_string(i));
            }
        }
        return results;
    }
};