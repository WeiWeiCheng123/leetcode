func removeElement(nums []int, val int) int {
    k := 0
    for _, s := range nums {
        if s != val {
            nums[k] = s
            k ++
        }
    }
    return k
}