class Solution {
public:
    int minimumLengthEncoding(vector<string>& words) {
	//方法1，反转+排序
        int N = words.size();
        // 反转每个单词
        vector<string> reversed_words;
        for (string word : words) {
            reverse(word.begin(), word.end());
            reversed_words.push_back(word);
        }
        // 字典序排序    
        sort(reversed_words.begin(), reversed_words.end());

        int res = 0;
        for (int i = 0; i < N; i++) {
            if (i+1 < N && reversed_words[i+1].find(reversed_words[i]) == 0) {
                // 当前单词是下一个单词的前缀，丢弃
            } else {
                res += reversed_words[i].length() + 1; // 单词加上一个 '#' 的长度
            }
        }
        return res;
    }
};
