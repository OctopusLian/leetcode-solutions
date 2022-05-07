class Solution:
    """
    @param source: 
    @param target: 
    @return: return the index
    """
    def strStr(self, source, target):
        for i in range(len(source) - len(target) + 1): #首字母出现在source中最后的位置
            index = i                                  #如果source长度为5，target长度为4
            flag = True                                #则超过1以后的就可以不匹配了
            for y in target:
                if source[index] == y:
                    index += 1
                else:
                    flag = False
                    break
            if flag:
                return i
        return -1