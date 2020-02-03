class Solution {
public:
    vector<vector<string>> partition(string s) {
        vector<vector<string> > result;
        if (s.empty()) return result;
        
        vector<string> palindromes;
        dfs(s, 0, palindromes, result);
        
        return result;
    }
private:
    void dfs(string s, int pos, vector<string> &palindromes, 
             vector<vector<string> > &ret) {
        
        if (pos == s.size()) {
            ret.push_back(palindromes);
            return;
        }
        
        for (int i = pos + 1; i <= s.size(); ++i) {
            string substr = s.substr(pos, i - pos);
            if (!isPalindrome(substr)) {
                continue;
            }
            
            palindromes.push_back(substr);
            dfs(s, i, palindromes, ret);
            palindromes.pop_back();
        }
    }
    
    bool isPalindrome(string s) {
        if (s.empty()) return false;
        
        int n = s.size();
        for (int i = 0; i < n; ++i) {
            if (s[i] != s[n - i - 1]) return false;
        }
        
        return true;
    }
};
