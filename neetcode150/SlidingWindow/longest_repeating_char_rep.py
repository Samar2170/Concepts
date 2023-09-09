class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        count={}
        i,j=0,0
        res=0
        maxCount=0
        while j < len(s):
            count[s[j]] = count.get(s[j],0) +1
            maxCount = max(maxCount,count[s[j]])
            if j-i+1-maxCount > k:
                count[s[i]] -=1
                i+=1
            res = max(res,j-i+1)
            j+=1
        return res

## create count dictionary, if length of window - maxCount > k, then we need to shrink the window else we can expand the window



if __name__ == "__main__":
    s = Solution()
    print(s.characterReplacement("ABAB",2))
    print(s.characterReplacement("AABABBA",1))
