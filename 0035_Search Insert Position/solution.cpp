class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        int n = nums.size();
        int start = 1;
        if (nums[n - 1] < target) {
            return n;
        }
        if (nums[0] >= target) {
            return 0;
        }
        if (nums[n / 2] < target) {
            start = n / 2;
        }
        for (int i = start; i < n; i++) {
            if (nums[i] >= target && nums[i - 1] <= target) {
                return i;
            }
        }
        return 0;
    }
};