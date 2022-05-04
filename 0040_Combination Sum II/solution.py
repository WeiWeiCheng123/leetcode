class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        combination = []  
        self.dfs(candidates, target, [], combination, 0)
        return combination
    
    def dfs(self, nums, target, path, combination, index):
        print(path)
        if target == 0:
            combination.append(path)
            return
        
        if target < 0:
            return
        
        for i in range(index, len(nums)):
            if i > index and nums[i] == nums[i - 1]:
                continue

            if nums[i] > target:
                break
                
            self.dfs(nums, target - nums[i], path + [nums[i]], combination, i + 1)