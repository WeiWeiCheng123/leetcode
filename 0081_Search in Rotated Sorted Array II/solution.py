class Solution:
    def search(self, nums: List[int], target: int) -> bool:
        left = 0
        right = len(nums) - 1
        while left <= right:
            if nums[left] != target and nums[right] != target:
                left += 1
                right += -1
            else :
                return True
        return False