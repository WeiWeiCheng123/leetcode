func threeSumClosest(nums []int, target int) int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    closest := nums[0] + nums[1] + nums[2]
    n := len(nums)
    for i, _ := range nums {
        left := i + 1
        right := n - 1
        for left < right {
            tmp := nums[i] + nums[left] + nums[right]
            if tmp == target{
                return target
            }
            if math.Abs(float64(tmp - target)) < math.Abs(float64(closest - target)) {
                closest = tmp
            }
            if tmp > target {
                right--
            } else {
                left++
            }
        }
    }
    return closest
}