class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        n = len(nums)
        start = 1
        if nums[n - 1] < target:
            return n
        if nums[0] >= target:
            return 0
        if nums[n // 2] < target:
            start = n // 2
        for i in range(start, n):
            if nums[i] >= target and nums[i - 1] <= target:
                return i
        return 0