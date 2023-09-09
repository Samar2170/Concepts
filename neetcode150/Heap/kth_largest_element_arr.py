from typing import List
import heapq

class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        nums = [-1*n for n in nums]

        heapq.heapify(nums)

        while k-1:
            heapq.heappop(nums)
            k-=1
        return -nums[0]

if __name__ == '__main__':
    nums = [3,2,1,5,6,4]
    k = 2
    print(Solution().findKthLargest(nums,k))