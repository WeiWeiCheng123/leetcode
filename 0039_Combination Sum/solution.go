func combinationSum(candidates []int, target int) [][]int {
    var combination [][]int
    var path []int
    dfs(candidates, target, path, &combination)
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
        path = append(path, nums[i])
        target += -nums[i]
        dfs(nums[i:], target, path, combination)
        path = path[:len(path)-1]
        target += nums[i]
    }
}