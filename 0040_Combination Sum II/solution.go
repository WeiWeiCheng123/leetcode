func combinationSum2(candidates []int, target int) [][]int {
    sort.Slice(candidates, func(i, j int) bool {
        return candidates[i] < candidates[j]
    })
    var combination [][]int
    dfs(candidates, target, make([]int, 0), &combination)
    return combination
}

func dfs(nums []int, target int, path []int, combination *[][]int) {
    if target == 0 {
        path1 := make([]int, len(path))
        copy(path1, path)
        *combination = append(*combination, path1)
        return
    }
    
    if target < 0 {
        return
    }
    
    for i := 0; i < len(nums); i++ {
        if nums[i] > target {
            break
        }
        path = append(path, nums[i])
        target += -nums[i]
        dfs(nums[i + 1:], target, path, combination)
        path = path[:len(path) - 1]
        target += nums[i]
        
        for i + 1 != len(nums) && nums[i + 1] == nums[i] {
            i++
        }
    }
}