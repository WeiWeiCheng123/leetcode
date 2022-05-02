class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        record = {}
        for i in range(len(nums)):
            if (target - nums[i] in record):
                return [i, record[target - nums[i]]]
            record[nums[i]] = i
        return None