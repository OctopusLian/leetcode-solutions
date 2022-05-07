class Solution {
public:
    /**
     * @param n: an integer
     * @return: an ineger f(n)
     */
    int fibonacci(int n) {
        std::vector<int>v={0,1};  
        while(n>v.size())  
            v.push_back(v.back()+v[v.size()-2]);  
        return v[n-1];  
    }
};