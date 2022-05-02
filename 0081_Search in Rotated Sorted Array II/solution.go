func search(nums []int, target int) bool {
    left := 0;
    right := len(nums) - 1
    for left <= right {
        if nums[left] != target && nums[right] != target {
            left ++
            right --
        } else {
            return true
        }
    }
    return false
}