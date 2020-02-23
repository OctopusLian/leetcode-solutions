class Solution {
public:
    string reverseLeftWords(string s, int n) {
        return s.substr(n, s.length() - 1) + s.substr(0,n);
    }
};
