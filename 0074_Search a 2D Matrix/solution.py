class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        for i in range(len(matrix)):
            left = 0
            right = len(matrix[i]) - 1
            while left <= right:
                if matrix[i][left] != target and matrix[i][right] != target:
                    left += 1
                    right += -1
                else:
                    return True
        return False
                