func twoSum(nums []int, target int) []int {
    record := make(map[int]int)
    for i, s := range nums {
        v, e := record[target - s]
        if e == true {
            output := []int{v, i}
            return output
        }
        record[s] = i
    }
    return nil
}