func threeSum(nums []int) [][]int {
    sort.Slice(nums, func(i, j int) bool { 
        return nums[i] < nums[j] 
    })
    var result [][]int
    for i := 0; i < len(nums); i++ {
        left := i + 1
        right := len(nums) - 1
        for left < right {
            if nums[i] + nums[left] + nums[right] > 0 {
                right--
            } else {
                if nums[i] + nums[left] + nums[right] < 0 {
                    left++
                } else {
                    triplet := []int{nums[i], nums[left], nums[right]}
                    result = append(result, triplet)
                    
                    for left < right && nums[left] == triplet[1] {
                        left++
                    }
                    for left < right && nums[right] == triplet[2] {
                        right--
                    }
                }
            }
        }
        for i + 1 < len(nums) && nums[i + 1] == nums[i] {
            i++
        }
        
    }
    return result
}