import heapq
from typing import List
class KthLargest:
    def __init__(self, k: int, nums: List[int]):
        self.k = k 
        self.minheap = nums
        heapq.heapify(self.minheap)
        
        while len(self.minheap) > self.k:
            heapq.heappop(self.minheap)
    def add(self, val: int) -> int:
        heapq.heappush(self.minheap,val)
        while len(self.minheap) > self.k:
            heapq.heappop(self.minheap)
        return self.minheap[0]

# 2. 215. Kth Largest Element in an Array
# https://leetcode.com/problems/kth-largest-element-in-an-array/
#
if __name__ == '__main__':
    nums = [4,5,8,2]
    k = 2
    kthLargest = KthLargest(k, nums)
    print(kthLargest.add(3))   # returns 4
    print(kthLargest.add(5))   # returns 5
    print(kthLargest.add(10))  # returns 5
    print(kthLargest.add(9))   # returns 8
    print(kthLargest.add(4))   # returns 8

