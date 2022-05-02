class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        result = []
        i = 0
        while i < len(nums):
            left = i + 1
            right = len(nums) - 1
            while left < right:
                if nums[i] + nums[left] + nums[right] > 0:
                    right += -1
                else:
                    if nums[i] + nums[left] + nums[right] < 0:
                        left += 1 
                    else:
                        triplet = [nums[i], nums[left], nums[right]]
                        result.append(triplet)
                        
                        while left < right and nums[left] == triplet[1]:
                            left += 1
                        while left < right and nums[right] == triplet[2]:
                            right += -1
            
            while i + 1 < len(nums) and nums[i + 1] == nums[i]:
                i += 1
            i += 1
        
        return result