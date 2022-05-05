func searchInsert(nums []int, target int) int {
    n := len(nums)
    start := 1
    if nums[n - 1] < target {
        return n
    }
    if nums[0] >= target {
        return 0
    }
    if nums[n / 2] < target {
        start = n / 2
    }
    for i := start; i < n; i++ {
        if nums[i] >= target && nums[i - 1] <= target {
            return i
        }
    }
    return 0
}