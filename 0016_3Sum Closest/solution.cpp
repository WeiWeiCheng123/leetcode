class Solution {
public:
    int threeSumClosest(vector<int>& nums, int target) {
        sort(nums.begin(), nums.end());
        int n = nums.size();
        int closest = nums[0] + nums[1] + nums[2];
        for (int i = 0; i < n-2; i++) {
            int left = i + 1;
            int right = n - 1;
            while (left < right) {
                int tmp = nums[i] + nums[left] + nums[right];
                if (tmp == target) {
                    return target;
                }
                if (abs(tmp - target) < abs(closest - target)) {
                    closest = tmp;
                }
                
                if (tmp > target) {
                    right += -1;
                }
                else {
                    left += 1;
                }
            }
        }
        return closest;
    }
};