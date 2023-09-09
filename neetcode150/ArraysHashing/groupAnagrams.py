from typing import List
class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        strset=set()
        result = {}
        for i in range(len(strs)):
            swrd = sorted(strs[i])
            swrd = ''.join(swrd)
            if swrd in strset:
                if swrd in result:
                    result[swrd].append(strs[i])
                else:
                    result[swrd] = [strs[i]]  
            else:
                strset.add(swrd)
                result[swrd]=[strs[i]]

        fres = []
        for key in result:
            fres.append(result[key])
        return fres



if __name__=='__main__':
    sol = Solution()
    print(sol.groupAnagrams(["eat","tea","tan","ate","nat","bat"]))
    print(sol.groupAnagrams(["pat","tap","eat","tea","tan","ate","nat","bat"]))


