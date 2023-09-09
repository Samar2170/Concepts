class Solution:
    def twoSum(self,nums,target):
        i=0
        j=len(nums)-1
        nums.sort()
        while i<j: 
            rem = target-nums[i] - nums[j]
            if rem==0:
                return [nums[i],nums[j]]
            elif rem>0:
                i+=1
            else:
                j-=1
        return []

if __name__ == "__main__":
    s = Solution()
    print(s.twoSum([2,7,11,15],9))
    print(s.twoSum([2,3,4],6))
    print(s.twoSum([-1,0,1,2,-1,-4],0))