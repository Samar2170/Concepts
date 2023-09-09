class Solution:
    def containsDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        return len(nums) != len(set(nums))

    def containsDuplicates2(self,nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        d = {}
        for i in nums:
            if i in d:
                return True
            else:
                d[i] = 1
        return False

if __name__ == "__main__":
    s = Solution()
    print(s.containsDuplicates([1,2,3,1]))
    print(s.containsDuplicates2([1,2,3,1]))