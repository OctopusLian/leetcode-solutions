class Solution(object):
    def groupAnagrams(self, strs):
        """
        :type strs: List[str]
        :rtype: List[List[str]]
        """
        classes = dict()
        for s in strs:
            sort_s = ''.join(sorted(s))
            classes[sort_s] = classes.get(sort_s, []) + [s]
        return [vals for vals in classes.values()]