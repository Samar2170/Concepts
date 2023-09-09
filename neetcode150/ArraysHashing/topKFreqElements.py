class Solution:
    def topKFrequent(self,nums, k):
        d={}
        for i in range(len(nums)):
            if nums[i] in d:
                d[nums[i]]+=1
            else:
                d[nums[i]]=1

        d = sorted(d.items(), key=lambda x: x[1], reverse=True)
        return [d[i][0] for i in range(k)]


if __name__ == "__main__":
    s = Solution()
    print(s.topKFrequent([1,1,1,2,2,3],2))
    print(s.topKFrequent([1],1))