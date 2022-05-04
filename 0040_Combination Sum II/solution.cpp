class Solution {
public:
    vector<vector<int>> combinationSum2(vector<int>& candidates, int target) {
        sort(candidates.begin(), candidates.end());
        vector<vector<int>> combination;
        vector<int> path;
        dfs(candidates, target, path, combination, 0);
        return combination;
    }
    
    void dfs(vector<int> nums, int target, vector<int> path, vector<vector<int>>& combination, int index) {
        if (target == 0) {
            combination.push_back(path);
            return;
        }
        
        if (target < 0) {
            return;
        }
        
        for (int i = index; i < nums.size(); i++) {
            if (nums[i] > target) {
                break;
            }
            target += -nums[i];
            path.push_back(nums[i]);
            dfs(nums, target, path, combination, i + 1);
            target += nums[i];
            path.pop_back();
            
            while (i + 1 < nums.size() && nums[i + 1] == nums[i]) {
                i++;
            }
        }
    }
};