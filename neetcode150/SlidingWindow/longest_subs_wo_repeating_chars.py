class Solution:
    def lengthOfLongestSubstring(self,s):
        if len(s)==0:
            return 0
        elif len(s)==1:
            return 1
        else:
            maxLen=0
            charMap={}
            currStr=''

            for i in range(len(s)):
                if s[i] not in charMap:
                    currStr+=s[i]
                    charMap[s[i]]=i
                else:
                    maxLen = max(maxLen,len(currStr))
                    currStr=s[charMap[s[i]]+1:i+1]
                    charMap[s[i]]=i

            
            maxLen = max(maxLen,len(currStr))

            return maxLen

    def longestSubstring(self,s):
        l,r=0,0
        charMap=set()
        res=0

        while r<len(s):
            while s[r] in charMap:
                charMap.remove(s[l])
                l+=1
            charMap.add(s[r])
            res = max(res,r-l+1) 
            r+=1
        print(charMap)
        return res

    # first iter cmap = {a} , r=0, l=0, res=1
    # second iter cmap = {a.b}, r=1, l=0, res=2
    # third iter cmap = {a,b,c}, r=2, l=0, res=3
    # fout iter cmap = {a,b,c}, r=3, l=1, res=3
    # fifth iter cmap = {a,b,c}, r=4, l=2, res=3
    


if __name__=='__main__':
    sol=Solution()
    # print(sol.lengthOfLongestSubstring("abcabcbb"))
    # print(sol.lengthOfLongestSubstring("bbbbb"))
    # print(sol.lengthOfLongestSubstring("pwwkew"))
    # print(sol.lengthOfLongestSubstring(" "))
    # print(sol.lengthOfLongestSubstring("au"))
    # print(sol.lengthOfLongestSubstring("dvdf"))
    # print(sol.lengthOfLongestSubstring("tmmzuxt")) # -> Exception
    # print(sol.lengthOfLongestSubstring("cdd"))

    print(sol.longestSubstring("abcabcbb")) 
    print(sol.longestSubstring("bbbbb"))
    print(sol.longestSubstring("pwwkew"))
    print(sol.longestSubstring(" "))
    print(sol.longestSubstring("au"))
    print(sol.longestSubstring("dvdf"))
    print(sol.longestSubstring("tmmzuxt")) # -> Exception
    print(sol.longestSubstring("cdd"))


## iterate over -> check if char in dict -> if present change currStr to indexed alt and change index in dict
#                                        -> else add as key to dict with index as value