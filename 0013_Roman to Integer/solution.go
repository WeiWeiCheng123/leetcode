func romanToInt(s string) int {
    roman := map[byte]int{
        73: 1,
        86: 5,
        88: 10,
        76: 50,
        67: 100,
        68: 500,
        77: 1000,
    }
    result := 0
    for i := 0; i < len(s); i++ {
        if i + 1 < len(s) && roman[s[i]] < roman[s[i + 1]] {
            result -= roman[s[i]]
            i++
        }
        result += roman[s[i]]
    }
    return result
}