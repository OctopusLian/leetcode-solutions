class Solution(object):
    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        if not len(matrix) or not len(matrix[0]): #特判一下矩阵为空的情况
            return False
        row = 0
        col = len(matrix[0]) - 1
        while row != len(matrix) and col != -1:
            if matrix[row][col] > target:
                col -= 1
            elif matrix[row][col] < target:
                row += 1   
            else:
                return True
        return False