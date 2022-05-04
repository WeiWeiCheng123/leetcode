class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        combination = []
        self.dfs(candidates, target, [], combination)
        return combination
    
    def dfs(self, nums, target, path, combination):
        if target == 0:
            combination.append(path)
            return
        
        if target < 0:
            return
        
        for i in range(len(nums)):
            self.dfs(nums[i:], target - nums[i], path + [nums[i]], combination)