func lengthOfLastWord(s string) int {
    count := 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] != 32 {
            count ++
        }
        if s[i] == 32 && count != 0 {
            return count
        }
    }
    return count
}