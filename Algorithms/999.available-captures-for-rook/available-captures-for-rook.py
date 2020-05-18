class Solution:
    def numRookCaptures(self, board: List[List[str]]) -> int:
        return (lambda x:x('pR')+x('Rp'))([''.join(board[x]+[' ']+[i[y] for i in board]).replace('.','') for x in range(8) for y in range(8) if board[x][y]=='R'][0].count)