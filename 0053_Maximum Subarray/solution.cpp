class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int maxSum = -1e5, result = -1e5;
        for (int i = 0; i < nums.size(); i++) {
            maxSum = max(nums[i], maxSum + nums[i]);
            result = max(result, maxSum);
        }
        return result;
    }
};