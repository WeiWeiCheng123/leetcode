class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int, int> m;
        vector<int> result;
        priority_queue<pair<int,int>> freq;
        for (int num: nums) {
            m[num]++;
        }
        for (auto it = m.begin(); it != m.end(); it++) {
            freq.push(make_pair(it->second, it->first));
        }
        while (k != 0) {
            result.push_back(freq.top().second);
            freq.pop();
            k--;
        }
        return result;
    }
};