class Solution:
    def validAnagram(self,s,t):
        d = {}
        for i in range(len(s)):
            if s[i] in d:
                d[s[i]]+=1

            else:
                d[s[i]]=1
        
        for i in range(len(t)):
            if t[i] in d:
                d[t[i]]-=1
            else:
                return False 

        if all(v==0 for v in d.values()):
            return True
        else:
            return False


if __name__ == "__main__":
    s = Solution()
    print(s.validAnagram("anagram","nagaram"))
    print(s.validAnagram("rat","car"))