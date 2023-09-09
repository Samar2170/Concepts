from re import L
from typing import List
class Solution:
    def search(self, nums: List[int], target: int) -> int:
        left,right = 0,len(nums)-1
        while left <=right:
            mid = (left+right)//2
            if nums[mid] == target:
                return mid
            elif nums[mid] < target:
                left = mid+1
            else:
                right = mid-1
        return -1
        
        
if __name__ == '__main__':
    s = Solution()
    print(s.search([1,2,3,4,5,6],3))
    print(s.search([1,2,3,4,5,6],6))
    print(s.search([1,2,3,4,5,6],1))
    print(s.search([1,2,3,4,5,6],7))
    print(s.search([1,2,3,4,5,6],0))
