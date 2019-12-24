class Solution(object):
    def deckRevealedIncreasing(self, deck):
        """
        :type deck: List[int]
        :rtype: List[int]
        """

        res = list()
        deck.sort(reverse = True)
        for card in deck:
            if len(res) >= 2:
                res = [res[-1]] + res[:-1]

            res.insert(0,card)

        return res
       
