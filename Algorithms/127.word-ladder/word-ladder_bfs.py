class Solution(object):
    def ladderLength(self, beginWord, endWord, wordList):
        """
        :type beginWord: str
        :type endWord: str
        :type wordList: List[str]
        :rtype: int
        """
        # init
        wordList = set(wordList)
        # BFS
        queue = collections.deque([(beginWord,1)])
        while queue:
            word,length = queue.popleft()
            if word == endWord:
                return length

            for i in range(len(word)):
                for c in 'abcedfghijklmnopqrstuvwxyz':
                    next_word = word[:i] + c + word[i+1:]
                    if next_word in wordList:
                        queue.append((next_word,length+1))
                        wordList.remove(next_word)

        return 0