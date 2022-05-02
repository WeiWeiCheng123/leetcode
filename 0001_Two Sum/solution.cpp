class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        map<int, int> record;
        for (int i = 0; i < nums.size(); i++) {
            if (record.find(target - nums[i]) != record.end())
                return {i, record[target - nums[i]]};
            record[nums[i]] = i;
        }
        return {};
    }
};