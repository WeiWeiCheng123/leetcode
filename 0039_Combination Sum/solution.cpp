class Solution {
public:
    vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
        vector<vector<int>> combination;
        vector<int> path;
        dfs(candidates, target, path, combination);
        return combination;
    }
    
    void dfs(vector<int> nums, int target, vector<int>& path, vector<vector<int>>& combination) {
        if (target == 0) {
            combination.push_back(path);
            return;
        }
        if (target < 0) {
            return;
        }
        for (int i = 0; i < nums.size(); i++) {
            target += -nums[i];
            path.push_back(nums[i]);
            dfs(vector<int>(nums.begin() + i, nums.end()), target, path, combination);
            path.pop_back();
            target += nums[i];
        }
    }
};