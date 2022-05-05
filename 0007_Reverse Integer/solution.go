func reverse(x int) int {
    var ans int
    for x != 0 {
        if ans > 0 && ans > (math.MaxInt32 - x % 10) / 10 {
            return 0
        }
        if ans < 0 && ans < (math.MinInt32 - x % 10) / 10 {
            return 0
        }
        ans = ans * 10 + x % 10
        x = x / 10
    }
    return ans
}