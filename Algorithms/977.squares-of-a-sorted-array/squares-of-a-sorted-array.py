class Solution:
    def sortedSquares(self, A: 'List[int]') -> 'List[int]':
        def f(x):
            return x*x
        return sorted(map(f,A))
