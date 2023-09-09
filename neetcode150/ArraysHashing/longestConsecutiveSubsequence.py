from typing import List
class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        nums.sort()
        currLen=1
        maxLen=0
        i=1
        while i < len(nums):
            # print(i, nums[i])
            if nums[i]-1==nums[i-1]:
                currLen+=1
            else:
                maxLen=max(currLen,maxLen)
                currLen=1
            maxLen=max(currLen,maxLen)
            i+=1
        return maxLen
        




if __name__=='__main__':
    sol = Solution()
    print(sol.longestConsecutive([100,4,200,1,3,2]))
    print(sol.longestConsecutive([0,3,7,2,5,8,4,6,0,1]))
    print(sol.longestConsecutive([0,3,7,2,5,8,4,6,0]))