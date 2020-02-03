class Solution(object):
    def wordBreak(self, s, wordDict):
        """
        :type s: str
        :type wordDict: List[str]
        :rtype: bool
        """
        if not s:
            return True
        if not wordDict:
            return False

        max_word_len = max([len(w) for w in wordDict])
        can_break = [True]
        for i in xrange(len(s)):
            can_break.append(False)
            for j in xrange(i, -1, -1):
                # optimize for too long interval
                if i - j + 1 > max_word_len:
                    break
                if can_break[j] and s[j:i + 1] in wordDict:
                    can_break[i + 1] = True
                    break
        return can_break[-1]
