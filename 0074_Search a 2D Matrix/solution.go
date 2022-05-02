func searchMatrix(matrix [][]int, target int) bool {
    for i, s := range matrix {
        left := 0
        right := len(s) - 1
        for left <= right {
            if matrix[i][left] != target && matrix[i][right] != target {
                left++
                right--
            } else {
                return true
            }
        }
    }
    return false
}