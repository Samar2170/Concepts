class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        if len(s1) > len(s2):
            return False
        strs = set()
        for s in s1:
            strs.add(s)
        i,j=0,0
        while i < len(s2):
            if s2[i] in strs:
                j = i
                while j < len(s2) and s2[j] in strs:
                    j += 1
                if j - i == len(s1):
                    return True
                i = j
            else:
                i += 1
        return False




if __name__=='__main__':
    sol = Solution()
    print(sol.checkInclusion("ab","eidbaooo"))
    print(sol.checkInclusion("ab","eidboaoo"))
    print(sol.checkInclusion("adc","dcda"))