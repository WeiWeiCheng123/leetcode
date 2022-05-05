class Solution:
    def reverse(self, x: int) -> int:
        ans = 0
        symbol = 1
        int_max = 2 ** 31 - 1
        int_min = -(2 ** 31)
        if x < 0:
            x = -x
            symbol = -1
        while x != 0:
            if ans > 0 and ans > (int_max - x % 10) // 10:
                return 0
            if ans < 0 and ans < (int_min - x % 10) // 10:
                return 0
            ans = ans * 10 + x % 10
            x = x // 10
            
        return ans * symbol