class Solution:
    """
    @param ipLines: ip  address
    @return: return highestFrequency ip address
    """
    def highestFrequency(self, ipLines):
        dict_num = {}
        for i in ipLines:
            if i in dict_num:
                dict_num[i] += 1
            else:
                dict_num[i] = 1

        return max(dict_num, key=lambda key: dict_num[key])