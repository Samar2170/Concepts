class Solution:
    def threeSum(self,nums,target=0):
        # import ipdb; ipdb.set_trace()
        i,j,k = 0,1,len(nums)-1

        nums.sort()
        res = []
        while i <len(nums)-2:
            if i>0 and nums[i]==nums[i-1]:
                i+=1
                continue
            j=i+1
            k=len(nums)-1
            while j<k:
                rem = target - nums[i] - nums[j] - nums[k]
                if rem == 0:
                    res.append([nums[i],nums[j],nums[k]])
                    j+=1
                    while nums[j-1]==nums[j] and j<k:
                        j+=1
                elif rem > 0:
                    j+=1
                else:
                    k-=1
            i+=1
        return res

if __name__ == "__main__":
    s = Solution()
    print(s.threeSum([-1,0,1,2,-1,-4],0))
    print(s.threeSum([2,7,11,15,0],9))
    print(s.threeSum([2,3,4],6))
    print(s.threeSum([0,0,0],0))


# nums[i]+nums[j]+nums[k] == target
# rem = target - nums[i] - nums[j] - nums[k] ; rem==0;