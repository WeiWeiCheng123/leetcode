func search(nums []int, target int) int {
    left := 0
    right := len(nums)-1
    for left <= right {
        if nums[left] != target && nums[right] != target {
            left++
            right--
        } else {
            if nums[left] == target {
                return left
            }
            return right
        }     
    }
    return -1
}