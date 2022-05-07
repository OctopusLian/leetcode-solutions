class Solution:
    """
    @param s: the article
    @return: the number of letters that are illegal
    """
    def count(self, s):
        st = True  #这个变量用来判断是否是句子开头
        count = 0
        for i in range(len(s)):
            c = s[i]
            if st and c>='a' and c<='z': #如果是句子开头，并且是小写字母，count+1
                count += 1
            if (c>='a' and c<='z') or (c >= 'A' and c <= 'Z'):#如果不是句子开头，就把这个变量置为false
                st = False
            if c == '.':#如果一个句子结束了，变量重新置为true
                st = True
            if (i>0 and (s[i-1]!='.' and s[i-1]!=',' and s[i-1]!=' ') and (c >= 'A' and c<='Z')):#非首字母 大写， count+1
                count += 1
        return count