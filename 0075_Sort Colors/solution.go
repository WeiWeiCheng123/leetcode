func sortColors(nums []int)  {
    tmp := 0
    n := len(nums)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if nums[i] > nums[j] {
                tmp = nums[i]
                nums[i] = nums[j]
                nums[j] = tmp
            }
        }
    }
}