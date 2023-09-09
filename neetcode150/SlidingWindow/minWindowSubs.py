class Solution:
    def minWindow(self, s: str, t: str) -> str:
        tmap={}
        for i in range(len(t)):
            if t[i] in tmap:
                tmap[t[i]]+=1
            else:
                tmap[t[i]]=1

        l=0
        res,reslen=[],float('inf')
        have,need=0,len(t)
        smap={}
        for i in range(len(s)):
            smap[s[i]] = 1 + smap.get(s[i],0)
            if s[i] in tmap and smap[s[i]]==tmap[s[i]]:
                have+=1
            while have==need:
                if (i-l+1) < reslen:
                    reslen = i-l+1
                    res = [l,i]
                smap[s[l]]-=1
                if s[l] in tmap and smap[s[l]]<tmap[s[l]]:
                    have-=1
                l+=1
        a,b=res
        return s[a:b+1] if reslen!=float('inf') else "" 



if __name__ == "__main__":
    s = Solution()
    print(s.minWindow("ADOBECODEBANC","ABC"))
