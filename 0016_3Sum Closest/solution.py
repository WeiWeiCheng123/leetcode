class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        closest = nums[0] + nums[1] + nums[2]
        nums.sort()
        n = len(nums)
        for i in range(n - 2):
            left = i + 1
            right = n - 1
            while left < right:
                tmp = nums[i] + nums[left] + nums[right]
                if tmp == target:
                    return target
                 
                if abs(tmp - target) < abs(closest - target):
                    closest = tmp
                    
                if tmp > target:
                    right += -1
                else:
                    left += 1
                    
        return closest