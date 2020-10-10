class Solution:
    def uniqueMorseRepresentations(self, words: List[str]) -> int:
        T = []
        moer = [".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
        str1 = 'abcdefghijklmnopqrstuvwxyz'
        zimu = []
        for i in str1:
            zimu.append(i)
        for i in range(len(words)):
            str2 = ''
            for j in range(len(words[i])):
                k = zimu.index(words[i][j])
                str2 += moer[k]
            if str2 not in T:
                T.append(str2)
        return len(T)